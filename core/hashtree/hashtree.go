package hashtree

import (
	"crypto/sha256"
	"errors"
	"io"
	"os"

	"github.com/cbergoon/merkletree"
)

// HashTreeContent implements the Content interface provided by merkletree
// and represents the content stored in the tree.
type HashTreeContent struct {
	x string
}

// CalculateHash hashes the values of a HashTreeContent
func (t HashTreeContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t HashTreeContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(HashTreeContent).x, nil
}

// NewHashTree build file to build hash tree
func NewHashTree(chunkPath []string) (*merkletree.MerkleTree, error) {
	if len(chunkPath) == 0 {
		return nil, errors.New("Empty data")
	}
	var list = make([]merkletree.Content, len(chunkPath))
	for i := 0; i < len(chunkPath); i++ {
		f, err := os.Open(chunkPath[i])
		if err != nil {
			return nil, err
		}
		temp, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		f.Close()
		list[i] = HashTreeContent{x: string(temp)}
	}

	//Create a new Merkle Tree from the list of Content
	return merkletree.NewTree(list)
}
