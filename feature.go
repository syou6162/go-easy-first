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

func extractFeature(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx)
	return features
}
