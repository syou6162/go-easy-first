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
