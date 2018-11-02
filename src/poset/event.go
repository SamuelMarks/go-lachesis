package poset

import (
	"crypto/ecdsa"
	"fmt"
	"reflect"

	"github.com/andrecronje/lachesis/src/crypto"
	"github.com/golang/protobuf/proto"
)

/*******************************************************************************
Comparison functions
*******************************************************************************/

func BytesEquals(this []byte, that []byte) bool {
	if len(this) != len(that) {
		return false
	}
	for i, v := range this {
		if v != that[i] {
			return false
		}
	}
	return true
}

func (this *BlockSignature) Equals(that *BlockSignature) bool {
	return reflect.DeepEqual(this.Validator, that.Validator) &&
		this.Index == that.Index &&
		this.Signature == that.Signature
}

func BlockSignatureListEquals(this []*BlockSignature, that []*BlockSignature) bool {
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

func (this *EventBody) Equals(that *EventBody) bool {
	return reflect.DeepEqual(this.Transactions, that.Transactions) &&
		reflect.DeepEqual(this.Parents, that.Parents) &&
		reflect.DeepEqual(this.Creator, that.Creator) &&
		this.Index == that.Index &&
		BlockSignatureListEquals(this.BlockSignatures, that.BlockSignatures) &&
		this.SelfParentIndex == that.SelfParentIndex &&
		this.OtherParentCreatorID == that.OtherParentCreatorID &&
		this.OtherParentIndex == that.OtherParentIndex &&
		this.CreatorID == that.CreatorID
}

func (this *EventCoordinates) Equals(that *EventCoordinates) bool {
	return this.Hash == that.Hash && this.Index == that.Index
}

func (this *Index) Equals(that *Index) bool {
	return this.ParticipantId == that.ParticipantId && this.Event.Equals(that.Event)
}

func IndexListEquals(this []*Index, that []*Index) bool {
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

func (this *Event) Equals(that *Event) bool {
	return this.Body.Equals(that.Body) &&
		this.Signature == that.Signature &&
		this.TopologicalIndex == that.TopologicalIndex &&
		this.Round == that.Round &&
		this.LamportTimestamp == that.LamportTimestamp &&
		this.RoundReceived == that.RoundReceived &&
		IndexListEquals(this.LastAncestors, that.LastAncestors) &&
		IndexListEquals(this.FirstDescendants, that.FirstDescendants) &&
		this.creator == that.creator &&
		reflect.DeepEqual(this.hash, that.hash) &&
		this.hex == that.hex &&
		BytesEquals(this.flagTable, that.flagTable)
}

/*******************************************************************************
EventBody
*******************************************************************************/

func (e *EventBody) Hash() ([]byte, error) {
	hashBytes, err := proto.Marshal(e)
	if err != nil {
		return nil, err
	}
	return crypto.SHA256(hashBytes), nil
}

/*******************************************************************************
Event
*******************************************************************************/

func GetIDIndex(o []*Index, id int64) int {
	for i, idx := range o {
		if idx.ParticipantId == id {
			return i
		}
	}

	return -1
}

func GetByID(o []*Index, id int64) (Index, bool) {
	for _, idx := range o {
		if idx.ParticipantId == id {
			return *idx, true
		}
	}

	return Index{}, false
}

// -----

// NewEvent creates new block event.
func NewEvent(transactions [][]byte, blockSignatures []*BlockSignature,
	parents []string, creator []byte, index int64,
	flagTable map[string]int64) Event {

	body := EventBody{
		Transactions:    transactions,
		BlockSignatures: blockSignatures,
		Parents:         parents,
		Creator:         creator,
		Index:           index,
	}

	// TODO: We shouldn't eat the error here...
	ft, _ := proto.Marshal(&FlagTableWrapper { Body: flagTable })

	return Event{
		Body:      &body,
		flagTable: ft,
		Round: -1,
		TopologicalIndex: -1,
		LamportTimestamp: -1,
		RoundReceived: -1,
	}
}

func (e *Event) SelfParent() string {
	return e.Body.Parents[0]
}

func (e *Event) OtherParent() string {
	return e.Body.Parents[1]
}

func (e *Event) Transactions() [][]byte {
	return e.Body.Transactions
}

func (e *Event) Index() int64 {
	return e.Body.Index
}

func (e *Event) BlockSignatures() []*BlockSignature {
	return e.Body.BlockSignatures
}

//True if Event contains a payload or is the initial Event of its creator
func (e *Event) IsLoaded() bool {
	if e.Body.Index == 0 {
		return true
	}

	hasTransactions := e.Body.Transactions != nil &&
		len(e.Body.Transactions) > 0

	return hasTransactions
}

//ecdsa sig
func (e *Event) Sign(privKey *ecdsa.PrivateKey) error {
	signBytes, err := e.Hash()
	if err != nil {
		return err
	}
	R, S, err := crypto.Sign(privKey, signBytes)
	if err != nil {
		return err
	}
	e.Signature = crypto.EncodeSignature(R, S)
	return err
}

func (e *Event) Verify() (bool, error) {
	pubBytes := e.Body.Creator
	pubKey := crypto.ToECDSAPub(pubBytes)

	signBytes, err := e.Hash()
	if err != nil {
		return false, err
	}

	r, s, err := crypto.DecodeSignature(e.Signature)
	if err != nil {
		return false, err
	}

	return crypto.Verify(pubKey, signBytes, r, s), nil
}

func (e *Event) Creator() string {
	if e.creator == "" {
		e.creator = fmt.Sprintf("0x%X", e.Body.Creator)
	}
	return e.creator
}

//sha256 hash of body
func (e *Event) Hash() ([]byte, error) {
	if len(e.hash) == 0 {
		hash, err := e.Body.Hash()
		if err != nil {
			return nil, err
		}
		e.hash = hash
	}
	return e.hash, nil
}

func (e *Event) Hex() string {
	if e.hex == "" {
		hash, _ := e.Hash()
		e.hex = fmt.Sprintf("0x%X", hash)
	}
	return e.hex
}

func (e *Event) SetRound(r int64) {
	e.Round = r
}

func (e *Event) SetLamportTimestamp(t int64) {
	e.LamportTimestamp = t
}

func (e *Event) SetRoundReceived(rr int64) {
	e.RoundReceived = rr
}

func (e *Event) SetWireInfo(SelfParentIndex,
	OtherParentCreatorID,
	OtherParentIndex,
	CreatorID int) {
	e.Body.SelfParentIndex = int64(SelfParentIndex)
	e.Body.OtherParentCreatorID = int64(OtherParentCreatorID)
	e.Body.OtherParentIndex = int64(OtherParentIndex)
	e.Body.CreatorID = int64(CreatorID)
}

func (e *Event) WireBlockSignatures() []WireBlockSignature {
	if e.Body.BlockSignatures != nil {
		wireSignatures := make([]WireBlockSignature, len(e.Body.BlockSignatures))
		for i, bs := range e.Body.BlockSignatures {
			wireSignatures[i] = bs.ToWire()
		}

		return wireSignatures
	}
	return nil
}

func (e *Event) ToWire() WireEvent {

	return WireEvent{
		Body: WireBody{
			Transactions:         e.Body.Transactions,
			SelfParentIndex:      int(e.Body.SelfParentIndex),
			OtherParentCreatorID: int(e.Body.OtherParentCreatorID),
			OtherParentIndex:     int(e.Body.OtherParentIndex),
			CreatorID:            int(e.Body.CreatorID),
			Index:                int(e.Body.Index),
			BlockSignatures:      e.WireBlockSignatures(),
		},
		Signature: e.Signature,
		FlagTable: e.flagTable,
	}
}

// GetFlagTable returns the flag table.
func (e *Event) GetUnmarshalledFlagTable() (result map[string]int64, err error) {
	wrapper := &FlagTableWrapper{}
	err = proto.Unmarshal(e.flagTable, wrapper)
	return wrapper.Body, err
}

// MargeFlagTable returns merged flag table object.
func (e *Event) MargeFlagTable(
	dst map[string]int64) (result map[string]int64, err error) {
	wrapper := &FlagTableWrapper{}
	if err := proto.Unmarshal(e.flagTable, wrapper); err != nil {
		return nil, err
	}

	for id, flag := range dst {
		if wrapper.Body[id] == 0 && flag == 1 {
			wrapper.Body[id] = 1
		}
	}
	return wrapper.Body, err
}

func rootSelfParent(participantID int) string {
	return fmt.Sprintf("Root%d", participantID)
}

/*******************************************************************************
Sorting
*******************************************************************************/

// ByTopologicalOrder implements sort.Interface for []Event based on
// the topologicalIndex field.
// THIS IS A PARTIAL ORDER
type ByTopologicalOrder []Event

func (a ByTopologicalOrder) Len() int      { return len(a) }
func (a ByTopologicalOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTopologicalOrder) Less(i, j int) bool {
	return a[i].TopologicalIndex < a[j].TopologicalIndex
}

// ByLamportTimestamp implements sort.Interface for []Event based on
// the lamportTimestamp field.
// THIS IS A TOTAL ORDER
type ByLamportTimestamp []Event

func (a ByLamportTimestamp) Len() int      { return len(a) }
func (a ByLamportTimestamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLamportTimestamp) Less(i, j int) bool {
	it, jt := -1, -1
	it = int(a[i].LamportTimestamp)
	jt = int(a[j].LamportTimestamp)
	if it != jt {
		return it < jt
	}

	wsi, _, _ := crypto.DecodeSignature(a[i].Signature)
	wsj, _, _ := crypto.DecodeSignature(a[j].Signature)
	return wsi.Cmp(wsj) < 0
}

/*******************************************************************************
 WireEvent
*******************************************************************************/

type WireBody struct {
	Transactions    [][]byte
	BlockSignatures []WireBlockSignature

	SelfParentIndex      int
	OtherParentCreatorID int
	OtherParentIndex     int
	CreatorID            int

	Index int
}

type WireEvent struct {
	Body      WireBody
	Signature string
	FlagTable []byte
}

func (we *WireEvent) BlockSignatures(validator []byte) []BlockSignature {
	if we.Body.BlockSignatures != nil {
		blockSignatures := make([]BlockSignature, len(we.Body.BlockSignatures))
		for k, bs := range we.Body.BlockSignatures {
			blockSignatures[k] = BlockSignature{
				Validator: validator,
				Index:     int64(bs.Index),
				Signature: bs.Signature,
			}
		}
		return blockSignatures
	}
	return nil
}
