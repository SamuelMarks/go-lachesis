// +build debug

// These functions are used only in debugging
package poset

import (
	"fmt"
)

func (p *Poset) PrintStat() {
	fmt.Println("****Known events:")
	for pid_id, index := range p.Store.KnownEvents() {
		fmt.Println("    index=", index, " peer=", p.Participants.ById[int64(pid_id)].NetAddr,
			" pubKeyHex=", p.Participants.ById[int64(pid_id)].PubKeyHex)
	}
}

// Gets output of events in a participant in json format
func (p *Poset) EventsJson(participantHex string) (string, error) {
	var result string
	result = ""

	// get participant Events with index > ct
	pevents, err := p.Store.ParticipantEvents(participantHex, -1)
	if err != nil {
		return result, err
	}
	for i, e := range pevents {
		ev, err := p.Store.GetEvent(e)
		if err != nil {
			return result, err
		}

		hash := ev.Hex()
		//Parents := ev.Body.Parents
		selfParent := ev.SelfParent()
		otherParent := ev.OtherParent()

		Creator := ev.Creator()
		//Index := ev.Body.Index
		//Signature := ev.Signature
		//TopologicalIndex: event.TopologicalIndex,
		//FlagTable: event.FlagTable,

		result += "{ hex: " + hash + ", parents: {" + selfParent + "," + otherParent + "}" + ", Creator: " + Creator + "}"
		if i < len(pevents) {
			result += ","
		}
	}
	return result, nil
}

// Gets events of a participant stored in this poset
func (p *Poset) GetEvents(phex string) ([]Event, error) {
	var events []Event
	// get participant Events with index > ct
	pevents, err := p.Store.ParticipantEvents(phex, -1)
	if err != nil {
		return []Event{}, err
	}
	for _, e := range pevents {
		ev, err := p.Store.GetEvent(e)
		if err != nil {
			return []Event{}, err
		}

		events = append(events, ev)
	}
	return events, nil
}

// Print the events
func (p *Poset) PrintEvents(events []Event) (string, error) {
	var result string
	result = ""
	for i, ev := range events {

		hash := ev.Hex()
		//Parents := ev.Body.Parents
		selfParent := ev.SelfParent()
		otherParent := ev.OtherParent()

		Creator := ev.Creator()
		//Index := ev.Body.Index
		//Signature := ev.Signature
		//TopologicalIndex: event.TopologicalIndex,
		//FlagTable: event.FlagTable,

		result += "{ hex: " + hash + ", parents: {" + selfParent + "," + otherParent + "}" + ", Creator: " + Creator + "}"
		if i < len(events) {
			result += ","
		}
	}

	return result, nil
}

func (s *BadgerStore) TopologicalEvents() ([]Event, error) {
	return s.dbTopologicalEvents()
}

// This is just a stub
func (s *InmemStore) TopologicalEvents() ([]Event, error) {
	return []Event{}, nil
}
