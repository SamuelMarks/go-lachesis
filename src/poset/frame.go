package poset

import (
	"github.com/andrecronje/lachesis/src/crypto"
	"github.com/golang/protobuf/proto"
)

func (f *Frame) Hash() ([]byte, error) {
	hashBytes, err := proto.Marshal(f)
	if err != nil {
		return nil, err
	}
	return crypto.SHA256(hashBytes), nil
}
