package main

import "fmt"

type Side int

const (
	LEFT Side = iota
	RIGHT
)

func addOneHandFeatures(features *[]string, w *Word, side Side) {
	*features = append(*features,
		fmt.Sprintf("side:%d+surface:%s", side, w.surface),
		fmt.Sprintf("side:%d+lemma:%s", side, w.lemma),
		fmt.Sprintf("side:%d+posTag:%s", side, w.posTag),
		fmt.Sprintf("side:%d+cposTag:%s", side, w.cposTag),
	)
}

func addLeftHandFeatures(features *[]string, state *State, idx int) {
	addOneHandFeatures(features, state.pending[idx], LEFT)
}

func addRightHandFeatures(features *[]string, state *State, idx int) {
	addOneHandFeatures(features, state.pending[idx+1], RIGHT)
}

func addBothHandFeatures(features *[]string, parent *Word, child *Word) {
	*features = append(*features,
		fmt.Sprintf("parent-surface:%s+child-surface:%s", parent.surface, child.surface),
		fmt.Sprintf("parent-lemma:%s+child-lemma:%s", parent.lemma, child.lemma),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s", parent.posTag, child.posTag),
		fmt.Sprintf("parent-cposTag:%s+child-cposTag:%s", parent.cposTag, child.cposTag),
	)
}

func extractAttachLeftFeatures(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx)
	parent := state.pending[idx]
	child := state.pending[idx+1]
	addBothHandFeatures(&features, parent, child)
	return features
}

func extractAttachRightFeatures(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx)
	parent := state.pending[idx+1]
	child := state.pending[idx]
	addBothHandFeatures(&features, parent, child)
	return features
}
