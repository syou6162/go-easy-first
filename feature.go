package main

import (
	"fmt"
	"math"
)

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

func distStr(dist int) string {
	d := "0"
	switch dist {
	case 1:
		d = "1"
	case 2:
		d = "2"
	case 3:
		d = "3"
	case 4:
		d = "4"
	default:
		d = "5"
	}
	return d
}

func addBothHandFeatures(features *[]string, parent *Word, child *Word) {
	dist := int(math.Abs(float64(parent.idx - child.idx)))

	*features = append(*features,
		fmt.Sprintf("parent-surface:%s+child-surface:%s", parent.surface, child.surface),
		fmt.Sprintf("parent-lemma:%s+child-lemma:%s", parent.lemma, child.lemma),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s", parent.posTag, child.posTag),
		fmt.Sprintf("parent-cposTag:%s+child-cposTag:%s", parent.cposTag, child.cposTag),
		fmt.Sprintf("dist:%s", distStr(dist)),
	)
}

func extractAttachLeftFeature(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx)
	parent := state.pending[idx]
	child := state.pending[idx+1]
	addBothHandFeatures(&features, parent, child)
	return features
}

func extractAttachRightFeature(state *State, idx int) []string {
	features := make([]string, 0)
	addLeftHandFeatures(&features, state, idx)
	addRightHandFeatures(&features, state, idx)
	parent := state.pending[idx+1]
	child := state.pending[idx]
	addBothHandFeatures(&features, parent, child)
	return features
}
