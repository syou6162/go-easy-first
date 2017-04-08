package main

import "fmt"

func addOneHandFeatures(features *[]string, w *Word, side string) {
	*features = append(*features,
		fmt.Sprintf("side:%s+surface:%s", side, w.surface),
		fmt.Sprintf("side:%s+lemma:%s", side, w.lemma),
		fmt.Sprintf("side:%s+posTag:%s", side, w.posTag),
		fmt.Sprintf("side:%s+cposTag:%s", side, w.cposTag),
	)
}

func addLeftHandFeatures(features *[]string, state *State, idx int) {
	addOneHandFeatures(features, state.pending[idx], "left")
}

func addRightHandFeatures(features *[]string, state *State, idx int) {
	addOneHandFeatures(features, state.pending[idx+1], "right")
}

func extractFeature(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx+1)
	return features
}
