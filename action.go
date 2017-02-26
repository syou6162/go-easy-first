package main

type StateAction func(state *State, idx int)

// AttachLeft は左側の単語を右側の単語の親にします
func AttachLeft(state *State, idx int) {
	parent := state.pending[idx]
	child := state.pending[idx+1]

	state.deletePending(idx + 1)
	parent.appendChild(child)
	state.arcs[child.idx] = parent.idx
}

// AttachRight は右側の単語を左側の単語の親にします
func AttachRight(state *State, idx int) {
	parent := state.pending[idx+1]
	child := state.pending[idx]

	state.deletePending(idx)
	parent.prependChild(child)
	state.arcs[child.idx] = parent.idx
}

// StateActions はActionの集合です
var StateActions = []StateAction{AttachLeft, AttachRight}

func main() {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 1, 2),
		makeWord("hang", "NNP", 2, 3),
		makeWord("plays", "VBZ", 3, 0),
		makeWord("elianti", "NNP", 4, 3),
		makeWord(".", ".", 5, 3),
	)
	s := &State{words, make(map[int]int)}
	AttachLeft(s, 3)
	AttachLeft(s, 3)
	AttachRight(s, 1)
	fmt.Println(s)
	for _, p := range s.pending {
		fmt.Println(p.children)
	}
	// fmt.Println("hgoe")
	// for _, p := range words {
	// 	fmt.Println(p)
	// }
}
