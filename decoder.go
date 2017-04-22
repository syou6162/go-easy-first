package main

import (
	"sort"
)

func isFinished(state *State) bool {
	return len(state.pending) == 1
}

func decode(weight *map[string]float64, state *State) {
	if isFinished(state) {
		// Do nothing
	} else {
		pair := BestActionIndexPair(weight, state)
		pair.action(state, pair.index)
		state.ResetFvCache(pair.index)
		decode(weight, state)
	}
}

func Decode(weight *map[string]float64, sent *Sentence) {
	tmp := make([]*Word, len(sent.words))
	copy(tmp, sent.words)
	s := NewState(sent.words)
	decode(weight, s)

	for child, parent := range s.arcs {
		sent.words[child].predHead = parent
	}
}

func beamSearch(weight *map[string]float64, beam States, beamWidth int) States {
	states := make(States, 0)
	for _, b := range beam {
		states = append(states, Expand(weight, b)...)
	}
	sort.Sort(states)
	return states[0:beamWidth]
}

func BeamSearch(weight *map[string]float64, sent *Sentence, beamWidth int) *State {
	tmp := make([]*Word, len(sent.words))
	copy(tmp, sent.words)
	beam := []State{*NewState(sent.words)}

	for {
		if isFinished(&beam[0]) {
			state := &beam[0]
			for child, parent := range state.arcs {
				sent.words[child].predHead = parent
			}
			return state
		}
		beam = beamSearch(weight, beam, beamWidth)
	}
}
