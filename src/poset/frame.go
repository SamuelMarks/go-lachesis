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

func RootListEquals(this []*Root, that []*Root) bool {
	if len(this) != len(that) {
		return false
	}
	for i, v := range this {
		if !v.Equals(that[i]) {
			return false
		}
	}
	return true
}

func EventListEquals(this []*Event, that []*Event) bool {
	if len(this) != len(that) {
		return false
	}
	for i, v := range this {
		if !v.Equals(that[i]) {
			return false
		}
	}
	return true
}

func (this *Frame) Equals(that *Frame) bool {
	return this.Round == that.Round &&
		RootListEquals(this.Roots, that.Roots) &&
		EventListEquals(this.Events, that.Events)
}
