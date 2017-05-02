package main

func isFinished(state *State) bool {
	return len(state.pending) == 1
}

func decode(weight *[]float64, state *State) {
	if isFinished(state) {
		// Do nothing
	} else {
		pair := BestActionIndexPair(weight, state)
		pair.action(state, pair.index)
		state.ResetFvCache(pair.index)
		decode(weight, state)
	}
}

func Decode(weight *[]float64, sent *Sentence) {
	s := NewState(sent.words)
	decode(weight, s)

	for child, parent := range s.arcs {
		sent.words[child].predHead = parent
	}
}
