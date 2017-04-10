package main

func isFinished(state *State) bool {
	return len(state.pending) == 1
}

func decode(weight *map[string]float64, state *State) {
	if isFinished(state) {
		// Do nothing
	} else {
		pair := BestActionIndexPair(weight, state)
		pair.action(state, pair.index)
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
