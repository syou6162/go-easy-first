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

func NilSafePendingWord(state *State, idx int) *Word {
	if idx < 0 || idx >= len(state.pending) {
		return nil
	} else {
		return state.pending[idx]
	}
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

func AddBigramFeatures(features *[]string, actName string, parent *Word, child *Word, prefix string) {
	if parent == nil || child == nil {
		return
	}

	plcp := NilSafePosTag(parent.LeftMostChild())
	prcp := NilSafePosTag(parent.RightMostChild())
	clcp := NilSafePosTag(child.LeftMostChild())
	crcp := NilSafePosTag(child.RightMostChild())
	dist := int(math.Abs(float64(parent.idx - child.idx)))

	*features = append(*features,
		fmt.Sprintf("%s+%s+parent-surface:%s+child-surface:%s", actName, prefix, parent.surface, child.surface),
		fmt.Sprintf("%s+%s+parent-surface:%s+child-posTag:%s", actName, prefix, parent.surface, child.posTag),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-surface:%s", actName, prefix, parent.posTag, child.surface),
		fmt.Sprintf("%s+%s+parent-lemma:%s+child-lemma:%s", actName, prefix, parent.lemma, child.lemma),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-posTag:%s", actName, prefix, parent.posTag, child.posTag),
		fmt.Sprintf("%s+%s+parent-cposTag:%s+child-cposTag:%s", actName, prefix, parent.cposTag, child.cposTag),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-posTag:%s+plcp:%s+prcp:%s", actName, prefix, parent.posTag, child.posTag, plcp, prcp),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-posTag:%s+plcp:%s+crcp:%s", actName, prefix, parent.posTag, child.posTag, plcp, crcp),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-posTag:%s+clcp:%s+prcp:%s", actName, prefix, parent.posTag, child.posTag, clcp, prcp),
		fmt.Sprintf("%s+%s+parent-posTag:%s+child-posTag:%s+clcp:%s+crcp:%s", actName, prefix, parent.posTag, child.posTag, clcp, crcp),
		fmt.Sprintf("%s+%s+dist:%s", actName, prefix, distStr(dist)),
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
	p0 := NilSafePendingWord(state, idx - 1)
	p1 := NilSafePendingWord(state, idx)
	p2 := NilSafePendingWord(state, idx + 1)
	p3 := NilSafePendingWord(state, idx + 2)
	AddBigramFeatures(&features, actName, p1, p2, "p_i+p_{i+1}")
	AddBigramFeatures(&features, actName, p1, p3, "p_i+p_{i+2}")
	AddBigramFeatures(&features, actName, p0, p1, "p_{i-1}+p_i")
	AddBigramFeatures(&features, actName, p0, p3, "p_{i-1}+p_{i+2}")
	AddBigramFeatures(&features, actName, p2, p3, "p_{i+1}+p_{i+2}")
	return features
}

func ExtractFeatures(state *State, pair ActionIndexPair) ([]string, error) {
	actName := runtime.FuncForPC(reflect.ValueOf(pair.action).Pointer()).Name()
	return extractFeatures(state, actName, pair.index), nil
}
