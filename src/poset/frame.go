package poset

import (
	"github.com/andrecronje/lachesis/src/crypto"
	"github.com/golang/protobuf/proto"
)

func (f *Frame) Marshal() ([]byte, error) {
	return proto.Marshal(f)
}

func (f *Frame) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, f)
}

func (f *Frame) Hash() ([]byte, error) {
	hashBytes, err := f.Marshal()
	if err != nil {
		return nil, err
	}
	return crypto.SHA256(hashBytes), nil
}
