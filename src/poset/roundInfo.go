package poset

type pendingRound struct {
	Index   int
	Decided bool
}

func NewRoundInfo() *RoundInfo {
	return &RoundInfo{
		Events: make(map[string]*RoundEvent),
	}
}

func (r *RoundInfo) AddEvent(x string, witness bool) {
	_, ok := r.Events[x]
	if !ok {
		r.Events[x] = &RoundEvent{
			Witness: witness,
		}
	}
}

func (r *RoundInfo) SetConsensusEvent(x string) {
	e, ok := r.Events[x]
	if !ok {
		e = &RoundEvent{}
	}
	e.Consensus = true
	r.Events[x] = e
}

func (r *RoundInfo) SetFame(x string, f bool) {
	e, ok := r.Events[x]
	if !ok {
		e = &RoundEvent{
			Witness: true,
		}
	}
	if f {
		e.Famous = Trilean_TRUE
	} else {
		e.Famous = Trilean_FALSE
	}
	r.Events[x] = e
}

//return true if no witnesses' fame is left undefined
func (r *RoundInfo) WitnessesDecided() bool {
	for _, e := range r.Events {
		if e.Witness && e.Famous == Trilean_UNDEFINED {
			return false
		}
	}
	return true
}

//return witnesses
func (r *RoundInfo) Witnesses() []string {
	var res []string
	for x, e := range r.Events {
		if e.Witness {
			res = append(res, x)
		}
	}
	return res
}

func (r *RoundInfo) RoundEvents() []string {
	var res []string
	for x, e := range r.Events {
		if !e.Consensus {
			res = append(res, x)
		}
	}
	return res
}

//return consensus events
func (r *RoundInfo) ConsensusEvents() []string {
	var res []string
	for x, e := range r.Events {
		if e.Consensus {
			res = append(res, x)
		}
	}
	return res
}

//return famous witnesses
func (r *RoundInfo) FamousWitnesses() []string {
	var res []string
	for x, e := range r.Events {
		if e.Witness && e.Famous == Trilean_TRUE {
			res = append(res, x)
		}
	}
	return res
}

func (r *RoundInfo) IsDecided(witness string) bool {
	w, ok := r.Events[witness]
	return ok && w.Witness && w.Famous != Trilean_UNDEFINED
}

func (r *RoundInfo) IsQueued() bool {
	return r.Queued
}
