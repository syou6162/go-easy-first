package main

type State struct {
	pending []*Word
	arcs    map[int]int
}

func NewState(pending []*Word) *State {
	p := make([]*Word, len(pending))
	copy(p, pending)
	return &State{p, make(map[int]int)}
}

func (state *State) deletePending(idx int) []*Word {
	state.pending = append(state.pending[:idx], state.pending[idx+1:]...)
	return state.pending
}
