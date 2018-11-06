package poset

import (
	"bytes"
	"encoding/json"

	"github.com/andrecronje/lachesis/src/crypto"
)

type Frame struct {
	Round  int64     //RoundReceived
	Roots  []Root  // [participant ID] => Root
	Events []Event //Event with RoundReceived = Round
}

//json encoding of Frame
func (f *Frame) Marshal() ([]byte, error) {

	var b bytes.Buffer
	enc := json.NewEncoder(&b) //will write to b
	if err := enc.Encode(f); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (f *Frame) Unmarshal(data []byte) error {

	b := bytes.NewBuffer(data)
	dec := json.NewDecoder(b) //will read from b
	err := dec.Decode(f)
	if err != nil {
		return err
	}
	for i, _ := range f.Events {
		f.Events[i].round = -1
		f.Events[i].roundReceived = -1
		f.Events[i].lamportTimestamp = -1
	}
	
	return nil
}

func (f *Frame) Hash() ([]byte, error) {
	hashBytes, err := f.Marshal()
	if err != nil {
		return nil, err
	}
	return crypto.SHA256(hashBytes), nil
}
