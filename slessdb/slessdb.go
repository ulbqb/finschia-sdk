package slessdb

import (
	"bytes"
	"crypto/sha256"

	"github.com/cosmos/iavl"
	dbm "github.com/tendermint/tm-db"
)

const (
	int64Size = 8
	hashSize  = sha256.Size
)

var (
	// nodeKeyFormat = iavl.NewKeyFormat('n', hashSize)  // n<hash>
	rootKeyFormat = iavl.NewKeyFormat('r', int64Size) // r<version>
)

type SLessDB struct {
	dbm.MemDB
	oracle *OracleClient
}

var _ dbm.DB = (*SLessDB)(nil)

// SlessDB is an stateless database backend using a MemDB and oracle for only iavl nodes.
//
// For stateless, all data is fethced via oracle. Other operatiosn are the same as MemDB's one.
// Support only Get and Has
func NewSlessDB(version int64, oracleClient OracleClientI, stores []string) *SLessDB {
	oracle := NewOracleClient(oracleClient)

	db := &SLessDB{
		MemDB:  *dbm.NewMemDB(),
		oracle: oracle,
	}

	// setup key stores
	empty := sha256.Sum256(nil)
	for i := range stores {
		rootHash := oracle.GetRootHash(stores[i])
		// An empty hash for the root hash means the store is empty at the version.
		// Empty stores doesn't store any root hash.
		if bytes.Equal(empty[:], rootHash) {
			continue
		}
		prefix := []byte("s/k:" + stores[i] + "/")
		key := append(prefix, rootKeyFormat.Key(version)...)
		db.Set(key, rootHash)
	}
	return db
}

func (db *SLessDB) Get(key []byte) ([]byte, error) {
	ok, store, hash := nodeHash(key)
	if !ok {
		return db.MemDB.Get(key)
	}

	node, _ := db.oracle.GetNode(string(store), hash)
	if node == nil {
		return nil, nil
	}
	var buf bytes.Buffer
	buf.Grow(node.encodedSize())

	if err := node.writeBytes(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (db *SLessDB) Has(key []byte) (bool, error) {
	ok, store, hash := nodeHash(key)
	if !ok {
		return db.MemDB.Has(key)
	}

	node, _ := db.oracle.GetNode(string(store), hash)

	return node != nil, nil
}

func nodeHash(key []byte) (bool, []byte, []byte) {
	storePrefix := []byte("s/k:")
	slash := []byte("/")
	nodePrefix := []byte("n")
	hashSize := sha256.Size

	if len(key) < len(storePrefix)+len(slash)+len(nodePrefix) {
		return false, nil, nil
	}

	// s/k:<STORE>/n<HASH>
	buf := append([]byte{}, key...)
	if !bytes.Equal(storePrefix, buf[:len(storePrefix)]) {
		return false, nil, nil
	}

	// <STORE>/n<HASH>
	buf = buf[len(storePrefix):]
	i := bytes.Index(buf, slash)
	if i <= 0 {
		return false, nil, nil
	}
	store := buf[:i]

	// n<HASH>
	buf = buf[i+len(slash):]
	if !bytes.Equal(nodePrefix, buf[:len(nodePrefix)]) {
		return false, nil, nil
	}

	// <HASH>
	hash := buf[1:]
	if len(hash) != hashSize {
		return false, nil, nil
	}

	return true, store, hash
}
