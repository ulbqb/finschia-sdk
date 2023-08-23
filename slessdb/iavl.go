package slessdb

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/Finschia/finschia-sdk/slessdb/internal/encoding"
)

type Node struct {
	key       []byte
	value     []byte
	hash      []byte
	leftHash  []byte
	rightHash []byte
	version   int64
	size      int64
	leftNode  *Node
	rightNode *Node
	height    int8
	persisted bool
}

func (node *Node) isLeaf() bool {
	return node.height == 0
}

func (node *Node) _hash() ([]byte, error) {
	if node.hash != nil {
		return node.hash, nil
	}

	h := sha256.New()
	buf := new(bytes.Buffer)
	if err := node.writeHashBytes(buf); err != nil {
		return nil, err
	}
	_, err := h.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	node.hash = h.Sum(nil)

	return node.hash, nil
}

func (node *Node) writeHashBytes(w io.Writer) error {
	err := encoding.EncodeVarint(w, int64(node.height))
	if err != nil {
		return errors.Wrap(err, "writing height")
	}
	err = encoding.EncodeVarint(w, node.size)
	if err != nil {
		return errors.Wrap(err, "writing size")
	}
	err = encoding.EncodeVarint(w, node.version)
	if err != nil {
		return errors.Wrap(err, "writing version")
	}

	// Key is not written for inner nodes, unlike writeBytes.

	if node.isLeaf() {
		err = encoding.EncodeBytes(w, node.key)
		if err != nil {
			return errors.Wrap(err, "writing key")
		}

		// Indirection needed to provide proofs without values.
		// (e.g. ProofLeafNode.ValueHash)
		valueHash := sha256.Sum256(node.value)

		err = encoding.EncodeBytes(w, valueHash[:])
		if err != nil {
			return errors.Wrap(err, "writing value")
		}
	} else {
		if node.leftHash == nil || node.rightHash == nil {
			return ErrEmptyChildHash
		}
		err = encoding.EncodeBytes(w, node.leftHash)
		if err != nil {
			return errors.Wrap(err, "writing left hash")
		}
		err = encoding.EncodeBytes(w, node.rightHash)
		if err != nil {
			return errors.Wrap(err, "writing right hash")
		}
	}

	return nil
}

// Writes the node as a serialized byte slice to the supplied io.Writer.
func (node *Node) writeBytes(w io.Writer) error {
	if node == nil {
		return errors.New("cannot write nil node")
	}
	cause := encoding.EncodeVarint(w, int64(node.height))
	if cause != nil {
		return errors.Wrap(cause, "writing height")
	}
	cause = encoding.EncodeVarint(w, node.size)
	if cause != nil {
		return errors.Wrap(cause, "writing size")
	}
	cause = encoding.EncodeVarint(w, node.version)
	if cause != nil {
		return errors.Wrap(cause, "writing version")
	}

	// Unlike writeHashBytes, key is written for inner nodes.
	cause = encoding.EncodeBytes(w, node.key)
	if cause != nil {
		return errors.Wrap(cause, "writing key")
	}

	if node.isLeaf() {
		cause = encoding.EncodeBytes(w, node.value)
		if cause != nil {
			return errors.Wrap(cause, "writing value")
		}
	} else {
		if node.leftHash == nil {
			return ErrLeftHashIsNil
		}
		cause = encoding.EncodeBytes(w, node.leftHash)
		if cause != nil {
			return errors.Wrap(cause, "writing left hash")
		}

		if node.rightHash == nil {
			return ErrRightHashIsNil
		}
		cause = encoding.EncodeBytes(w, node.rightHash)
		if cause != nil {
			return errors.Wrap(cause, "writing right hash")
		}
	}
	return nil
}

var (
	ErrCloneLeafNode  = fmt.Errorf("attempt to copy a leaf node")
	ErrEmptyChildHash = fmt.Errorf("found an empty child hash")
	ErrLeftHashIsNil  = fmt.Errorf("node.leftHash was nil in writeBytes")
	ErrRightHashIsNil = fmt.Errorf("node.rightHash was nil in writeBytes")
)

func (node *Node) encodedSize() int {
	n := 1 +
		encoding.EncodeVarintSize(node.size) +
		encoding.EncodeVarintSize(node.version) +
		encoding.EncodeBytesSize(node.key)
	if node.isLeaf() {
		n += encoding.EncodeBytesSize(node.value)
	} else {
		n += encoding.EncodeBytesSize(node.leftHash) +
			encoding.EncodeBytesSize(node.rightHash)
	}
	return n
}
