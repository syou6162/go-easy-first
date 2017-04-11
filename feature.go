package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func NilSafePosTag(w *Word) string {
	posTag := ""
	if w != nil {
		posTag = w.posTag
	}
	return posTag
}

func addOneHandFeatures(features *[]string, state *State, idx int, prefix string) {
	if idx < 0 || idx >= len(state.pending) {
		return
	}
	w := state.pending[idx]
	lcp := NilSafePosTag(w.LeftMostChild())
	rcp := NilSafePosTag(w.RightMostChild())
	*features = append(*features,
		fmt.Sprintf("%s+surface:%s", prefix, w.surface),
		fmt.Sprintf("%s+lemma:%s", prefix, w.lemma),
		fmt.Sprintf("%s+posTag:%s", prefix, w.posTag),
		fmt.Sprintf("%s+cposTag:%s", prefix, w.cposTag),
		fmt.Sprintf("%s+posTag:%s+leftmost:%s", prefix, w.posTag, lcp),
		fmt.Sprintf("%s+posTag:%s+rightmost:%s", prefix, w.posTag, rcp),
		fmt.Sprintf("%s+posTag:%s+leftmost:%s+rightmost:%s", prefix, w.posTag, lcp, rcp),
	)
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
	plcp := NilSafePosTag(parent.LeftMostChild())
	prcp := NilSafePosTag(parent.RightMostChild())
	clcp := NilSafePosTag(child.LeftMostChild())
	crcp := NilSafePosTag(child.RightMostChild())
	dist := int(math.Abs(float64(parent.idx - child.idx)))

	*features = append(*features,
		fmt.Sprintf("parent-surface:%s+child-surface:%s", parent.surface, child.surface),
		fmt.Sprintf("parent-surface:%s+child-posTag:%s", parent.surface, child.posTag),
		fmt.Sprintf("parent-posTag:%s+child-surface:%s", parent.posTag, child.surface),
		fmt.Sprintf("parent-lemma:%s+child-lemma:%s", parent.lemma, child.lemma),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s", parent.posTag, child.posTag),
		fmt.Sprintf("parent-cposTag:%s+child-cposTag:%s", parent.cposTag, child.cposTag),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s+plcp:%s+prcp:%s", parent.posTag, child.posTag, plcp, prcp),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s+plcp:%s+crcp:%s", parent.posTag, child.posTag, plcp, crcp),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s+clcp:%s+prcp:%s", parent.posTag, child.posTag, clcp, prcp),
		fmt.Sprintf("parent-posTag:%s+child-posTag:%s+clcp:%s+crcp:%s", parent.posTag, child.posTag, clcp, crcp),
		fmt.Sprintf("dist:%s", distStr(dist)),
	)
}

func extractAttachLeftFeatures(state *State, idx int) []string {
	features := make([]string, 0)
	addOneHandFeatures(&features, state, idx-2, "p_i-2")
	addOneHandFeatures(&features, state, idx-1, "p_i-1")
	addOneHandFeatures(&features, state, idx, "p_i")
	addOneHandFeatures(&features, state, idx+1, "p_i+1")
	addOneHandFeatures(&features, state, idx+2, "p_i+2")
	addOneHandFeatures(&features, state, idx+3, "p_i+3")
	parent := state.pending[idx]
	child := state.pending[idx+1]
	addBothHandFeatures(&features, parent, child)
	return features
}

func extractAttachRightFeatures(state *State, idx int) []string {
	features := make([]string, 0)
	addOneHandFeatures(&features, state, idx-2, "p_i-2")
	addOneHandFeatures(&features, state, idx-1, "p_i-1")
	addOneHandFeatures(&features, state, idx, "p_i")
	addOneHandFeatures(&features, state, idx+1, "p_i+1")
	addOneHandFeatures(&features, state, idx+2, "p_i+2")
	addOneHandFeatures(&features, state, idx+3, "p_i+3")
	parent := state.pending[idx+1]
	child := state.pending[idx]
	addBothHandFeatures(&features, parent, child)
	return features
}

func ExtractFeatures(state *State, pair ActionIndexPair) ([]string, error) {
	if reflect.ValueOf(pair.action).Pointer() == reflect.ValueOf(AttachLeft).Pointer() {
		return extractAttachLeftFeatures(state, pair.index), nil
	} else if reflect.ValueOf(pair.action).Pointer() == reflect.ValueOf(AttachRight).Pointer() {
		return extractAttachRightFeatures(state, pair.index), nil
	} else {
		return nil, errors.New("wrong action")
	}
}
