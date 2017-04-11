package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func NilSafePosTag(w *Word) string {
	posTag := ""
	if w != nil {
		posTag = w.posTag
	}
	return posTag
}

func addUnigramFeatures(features *[]string, state *State, actName string, idx int, prefix string) {
	if idx < 0 || idx >= len(state.pending) {
		return
	}
	w := state.pending[idx]
	lcp := NilSafePosTag(w.LeftMostChild())
	rcp := NilSafePosTag(w.RightMostChild())
	*features = append(*features,
		fmt.Sprintf("%s+%s+surface:%s", actName, prefix, w.surface),
		fmt.Sprintf("%s+%s+lemma:%s", actName, prefix, w.lemma),
		fmt.Sprintf("%s+%s+posTag:%s", actName, prefix, w.posTag),
		fmt.Sprintf("%s+%s+cposTag:%s", actName, prefix, w.cposTag),
		fmt.Sprintf("%s+%s+posTag:%s+leftmost:%s", actName, prefix, w.posTag, lcp),
		fmt.Sprintf("%s+%s+posTag:%s+rightmost:%s", actName, prefix, w.posTag, rcp),
		fmt.Sprintf("%s+%s+posTag:%s+leftmost:%s+rightmost:%s", actName, prefix, w.posTag, lcp, rcp),
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

func AddBigramFeatures(features *[]string, actName string, parent *Word, child *Word) {
	if parent == nil || child == nil {
		return
	}

	plcp := NilSafePosTag(parent.LeftMostChild())
	prcp := NilSafePosTag(parent.RightMostChild())
	clcp := NilSafePosTag(child.LeftMostChild())
	crcp := NilSafePosTag(child.RightMostChild())
	dist := int(math.Abs(float64(parent.idx - child.idx)))

	*features = append(*features,
		fmt.Sprintf("%s+parent-surface:%s+child-surface:%s", actName, parent.surface, child.surface),
		fmt.Sprintf("%s+parent-surface:%s+child-posTag:%s", actName, parent.surface, child.posTag),
		fmt.Sprintf("%s+parent-posTag:%s+child-surface:%s", actName, parent.posTag, child.surface),
		fmt.Sprintf("%s+parent-lemma:%s+child-lemma:%s", actName, parent.lemma, child.lemma),
		fmt.Sprintf("%s+parent-posTag:%s+child-posTag:%s", actName, parent.posTag, child.posTag),
		fmt.Sprintf("%s+parent-cposTag:%s+child-cposTag:%s", actName, parent.cposTag, child.cposTag),
		fmt.Sprintf("%s+parent-posTag:%s+child-posTag:%s+plcp:%s+prcp:%s", actName, parent.posTag, child.posTag, plcp, prcp),
		fmt.Sprintf("%s+parent-posTag:%s+child-posTag:%s+plcp:%s+crcp:%s", actName, parent.posTag, child.posTag, plcp, crcp),
		fmt.Sprintf("%s+parent-posTag:%s+child-posTag:%s+clcp:%s+prcp:%s", actName, parent.posTag, child.posTag, clcp, prcp),
		fmt.Sprintf("%s+parent-posTag:%s+child-posTag:%s+clcp:%s+crcp:%s", actName, parent.posTag, child.posTag, clcp, crcp),
		fmt.Sprintf("dist:%s", distStr(dist)),
	)
}

func AddUnigramFeatures(features *[]string, state *State, actName string, idx int) {
	addUnigramFeatures(features, state, actName, idx-2, "p_i-2")
	addUnigramFeatures(features, state, actName, idx-1, "p_i-1")
	addUnigramFeatures(features, state, actName, idx, "p_i")
	addUnigramFeatures(features, state, actName, idx+1, "p_i+1")
	addUnigramFeatures(features, state, actName, idx+2, "p_i+2")
	addUnigramFeatures(features, state, actName, idx+3, "p_i+3")
}

func extractFeatures(state *State, actName string, idx int) []string {
	features := make([]string, 0)
	AddUnigramFeatures(&features, state, actName, idx)
	parent := state.pending[idx]
	child := state.pending[idx+1]
	AddBigramFeatures(&features, actName, parent, child)
	return features
}

func ExtractFeatures(state *State, pair ActionIndexPair) ([]string, error) {
	actName := runtime.FuncForPC(reflect.ValueOf(pair.action).Pointer()).Name()
	return extractFeatures(state, actName, pair.index), nil
}
