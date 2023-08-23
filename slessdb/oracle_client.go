package slessdb

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"

	ics23 "github.com/confio/ics23/go"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/merkle"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
)

// OracleClient Interface
type OracleClientI interface {
	Get([]byte) []byte
}

// OracleClient
type OracleClient struct {
	oracle      OracleClientI
	accessedKey AccessedKey
}

func NewOracleClient(oracle OracleClientI) *OracleClient {
	return &OracleClient{
		oracle:      oracle,
		accessedKey: AccessedKey{},
	}
}

func (c *OracleClient) query(path, data string) *resultABCIQuery {
	b := c.oracle.Get([]byte(fmt.Sprintf("abci_query?path=%s&data=%s", url.PathEscape(path), data)))
	result := resultABCIQuery{}
	if err := tmjson.Unmarshal(b, &result); err != nil {
		panic(err)
	}
	return &result
}

func (c *OracleClient) getProof(path, data string) ([]*ics23.CommitmentProof, bool) {
	// if c.accessedKey.Has(path, data) {
	// 	return nil, true
	// }
	// c.accessedKey.Set(path, data)

	q := c.query("store/"+path, data).Response

	// cannot get node from empty store
	if q.ProofOps == nil {
		return []*ics23.CommitmentProof{}, false
	}

	ops := make([]*crypto.ProofOp, len(q.ProofOps.Ops))
	for i := range q.ProofOps.Ops {
		op := q.ProofOps.Ops[i]
		ops[i] = &op
	}

	cops, err := convertToCommitmentProofs(ops)
	if err != nil {
		panic(err)
	}
	return cops, false
}

func (c *OracleClient) GetRootHash(store string) []byte {
	proofs, _ := c.getProof(keyPath(store), dataString([]byte("roothash")))
	rootHash := proofs[len(proofs)-1].GetExist().Value
	return rootHash
}

func (c *OracleClient) GetNode(store string, hash []byte) (*Node, bool) {
	proofs, accessed := c.getProof(nodePath(store), dataString(hash))
	if accessed {
		return nil, accessed
	}

	if len(proofs) == 0 {
		return nil, false
	}

	proof := proofs[0]
	var node *Node

	ep := getExistenceProof(proof)[0]

	leaf, err := fromLeafOp(ep.GetLeaf(), ep.Key, ep.Value)
	if err != nil {
		panic(err)
	}
	if bytes.Equal(hash, leaf.hash) {
		node = leaf
	}

	if node == nil {
		prevHash := leaf.hash
		path := ep.GetPath()
		for j := range path {
			inner, err := fromInnerOp(path[j], prevHash)
			if err != nil {
				panic(err)
			}
			if bytes.Equal(hash, inner.hash) {
				node = inner
				break
			}
			prevHash = inner.hash
		}
	}

	if node == nil {
		panic("something wrong")
	} else {
		node.key = leaf.key
	}

	// TODO: this is work around so it need to check if it able to decode node value to empty byte array
	if node.isLeaf() && node.value == nil {
		node.value = []byte{}
	}

	return node, false
}

func keyPath(store string) string {
	return fmt.Sprintf("%s/key", store)
}

func nodePath(store string) string {
	return fmt.Sprintf("%s/node", store)
}

// CommitmentOp
type CommitmentOp struct {
	Type  string
	Spec  *ics23.ProofSpec
	Key   []byte
	Proof *ics23.CommitmentProof
}

var _ merkle.ProofOperator = CommitmentOp{}

func (op CommitmentOp) GetKey() []byte {
	return op.Key
}

func (op CommitmentOp) ProofOp() crypto.ProofOp {
	bz, err := op.Proof.Marshal()
	if err != nil {
		panic(err.Error())
	}
	return crypto.ProofOp{
		Type: op.Type,
		Key:  op.Key,
		Data: bz,
	}
}

func (op CommitmentOp) Run(args [][]byte) ([][]byte, error) {
	// calculate root from proof
	root, err := op.Proof.Calculate()
	if err != nil {
		return nil, fmt.Errorf("could not calculate root for proof: %v", err)
	}
	// Only support an existence proof or nonexistence proof (batch proofs currently unsupported)
	switch len(args) {
	case 0:
		// Args are nil, so we verify the absence of the key.
		absent := ics23.VerifyNonMembership(op.Spec, root, op.Proof, op.Key)
		if !absent {
			return nil, fmt.Errorf("proof did not verify absence of key: %s", string(op.Key))
		}

	case 1:
		// Args is length 1, verify existence of key with value args[0]
		if !ics23.VerifyMembership(op.Spec, root, op.Proof, op.Key, args[0]) {
			return nil, fmt.Errorf("proof did not verify existence of key %s with given value %x", op.Key, args[0])
		}
	default:
		return nil, fmt.Errorf("args must be length 0 or 1, got: %d", len(args))
	}

	return [][]byte{root}, nil
}

// NodeSet
type NodeSet map[string]*Node

func (s NodeSet) Set(n *Node) {
	key := fmt.Sprintf("%x", n.hash)
	if s[key] == nil {
		s[key] = n
	}
}

func (s NodeSet) Get(hash []byte) *Node {
	return s[fmt.Sprintf("%x", hash)]
}

func (s NodeSet) Has(hash []byte) bool {
	return s.Get(hash) != nil
}

func (s NodeSet) List() []*Node {
	ns := []*Node{}
	for i := range s {
		ns = append(ns, s[i])
	}
	return ns
}

// AccessedKey
type AccessedKey map[string]map[string]struct{}

func (ak AccessedKey) Set(k1, k2 string) {
	_, ok := ak[k1]
	if !ok {
		ak[k1] = map[string]struct{}{}
	}
	ak[k1][k2] = struct{}{}
}

func (ak AccessedKey) Has(k1, k2 string) bool {
	_, ok := ak[k1]
	if !ok {
		return false
	}
	_, ok = ak[k1][k2]
	return ok
}

// helper
func getExistenceProof(cp *ics23.CommitmentProof) []*ics23.ExistenceProof {
	eps := []*ics23.ExistenceProof{}
	switch cp.GetProof().(type) {
	case *ics23.CommitmentProof_Exist:
		eps = append(eps, cp.GetExist())
	case *ics23.CommitmentProof_Nonexist:
		nep := cp.GetNonexist()
		if nep.Left != nil {
			eps = append(eps, nep.Left)
		}
		if nep.Right != nil {
			eps = append(eps, nep.Right)
		}
	}
	return eps
}

func convertToCommitmentProofs(ops []*crypto.ProofOp) ([]*ics23.CommitmentProof, error) {
	cps := make([]*ics23.CommitmentProof, 0)
	for _, op := range ops {
		op, err := commitmentOpDecoder(*op)
		if err != nil {
			return nil, err
		}
		commitmentOp := op.(CommitmentOp)
		commitmentProof := commitmentOp.Proof
		if err != nil {
			return nil, err
		}
		cps = append(cps, commitmentProof)
	}
	return cps, nil
}

func commitmentOpDecoder(pop crypto.ProofOp) (merkle.ProofOperator, error) {
	var spec *ics23.ProofSpec
	switch pop.Type {
	case "ics23:iavl":
		spec = ics23.IavlSpec
	case "ics23:simple":
		spec = ics23.TendermintSpec
	default:
		return nil, fmt.Errorf("unexpected ProofOp.Type; got %s, want supported ics23 subtypes 'ProofOpIAVLCommitment' or 'ProofOpSimpleMerkleCommitment'", pop.Type)
	}

	proof := &ics23.CommitmentProof{}
	err := proof.Unmarshal(pop.Data)
	if err != nil {
		return nil, err
	}

	op := CommitmentOp{
		Type:  pop.Type,
		Key:   pop.Key,
		Spec:  spec,
		Proof: proof,
	}
	return op, nil
}

func dataString(b []byte) string {
	return hex.EncodeToString(b)
}

type resultABCIQuery struct {
	Response abci.ResponseQuery `json:"response"`
}

const (
	// lengthByte is the length prefix prepended to each of the sha256 sub-hashes
	lengthByte byte = 0x20
)

func fromInnerOp(iop *ics23.InnerOp, prevHash []byte) (*Node, error) {
	r := bytes.NewReader(iop.Prefix)
	height, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}
	size, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}
	version, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}

	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	if b != lengthByte {
		return nil, errors.New("expected length byte (0x20")
	}
	var left, right []byte
	// if left is empty, skip to right
	if r.Len() != 0 {
		left = make([]byte, lengthByte)
		n, err := r.Read(left)
		if err != nil {
			return nil, err
		}
		if n != 32 {
			return nil, errors.New("couldn't read left hash")
		}
		b, err = r.ReadByte()
		if err != nil {
			return nil, err
		}
		if b != lengthByte {
			return nil, errors.New("expected length byte (0x20")
		}
	}

	if len(iop.Suffix) > 0 {
		right = make([]byte, lengthByte)
		r = bytes.NewReader(iop.Suffix)
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		if b != lengthByte {
			return nil, errors.New("expected length byte (0x20")
		}

		n, err := r.Read(right)
		if err != nil {
			return nil, err
		}
		if n != 32 {
			return nil, errors.New("couldn't read right hash")
		}
	}

	if left == nil {
		left = prevHash
	} else if right == nil {
		right = prevHash
	}

	node := &Node{
		leftHash:  left,
		rightHash: right,
		version:   version,
		size:      size,
		height:    int8(height),
	}

	_, err = node._hash()
	if err != nil {
		return nil, err
	}

	return node, nil
}

func fromLeafOp(lop *ics23.LeafOp, key, value []byte) (*Node, error) {
	r := bytes.NewReader(lop.Prefix)
	height, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}
	if height != 0 {
		return nil, errors.New("height should be 0 in the leaf")
	}
	size, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}
	if size != 1 {
		return nil, errors.New("size should be 1 in the leaf")
	}
	version, err := binary.ReadVarint(r)
	if err != nil {
		return nil, err
	}
	node := &Node{
		key:     key,
		value:   value,
		size:    size,
		version: version,
	}

	_, err = node._hash()
	if err != nil {
		return nil, err
	}

	return node, nil
}
