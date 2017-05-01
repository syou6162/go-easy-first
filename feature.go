package main

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
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
		actName+"+"+prefix+"+surface:"+w.surface,
		actName+"+"+prefix+"+lemma:"+w.lemma,
		actName+"+"+prefix+"+posTag:"+w.posTag,
		actName+"+"+prefix+"+cposTag:"+w.cposTag,
		actName+"+"+prefix+"+posTag:"+w.posTag+"+leftmost:"+lcp,
		actName+"+"+prefix+"+posTag:"+w.posTag+"+rightmost:"+rcp,
		actName+"+"+prefix+"+posTag:"+w.posTag+"+leftmost:"+lcp+"+rightmost:"+rcp,
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

	*features = append(*features,
		actName+"+"+prefix+"+parent-surface:"+parent.surface+"+child-surface:"+child.surface,
		actName+"+"+prefix+"+parent-surface:"+parent.surface+"+child-posTag:"+child.posTag,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-surface:"+child.surface,
		actName+"+"+prefix+"+parent-lemma:"+parent.lemma+"+child-lemma:"+child.lemma,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-posTag:"+child.posTag,
		actName+"+"+prefix+"+parent-cposTag:"+parent.cposTag+"+child-cposTag:"+child.cposTag,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-posTag:"+child.posTag+"+plcp:"+plcp+"+prcp:"+prcp,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-posTag:"+child.posTag+"+plcp:"+plcp+"+crcp:"+crcp,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-posTag:"+child.posTag+"+clcp:"+clcp+"+prcp:"+prcp,
		actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-posTag:"+child.posTag+"+clcp:"+clcp+"+crcp:"+crcp,
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

func hasNoChildren(w *Word) bool {
	return len(w.children) == 0
}

func addStructuralSingleFeatures(features *[]string, state *State, actName string, idx int, prefix string) {
	if idx < 0 || idx >= len(state.pending) {
		return
	}
	w := state.pending[idx]
	*features = append(*features,
		actName+"+"+prefix+"+len:"+strconv.Itoa(len(w.children)),
		actName+"+"+prefix+"+no-children:"+strconv.FormatBool(hasNoChildren(w)),
	)
}

func AddStructuralSingleFeatures(features *[]string, state *State, actName string, idx int) {
	addStructuralSingleFeatures(features, state, actName, idx-2, "p_i-2")
	addStructuralSingleFeatures(features, state, actName, idx-1, "p_i-1")
	addStructuralSingleFeatures(features, state, actName, idx, "p_i")
	addStructuralSingleFeatures(features, state, actName, idx+1, "p_i+1")
	addStructuralSingleFeatures(features, state, actName, idx+2, "p_i+2")
	addStructuralSingleFeatures(features, state, actName, idx+3, "p_i+3")
}

func addStructuralPairFeatures(features *[]string, actName string, left *Word, right *Word, prefix string) {
	if left == nil || right == nil {
		return
	}
	dist := int(math.Abs(float64(left.idx - right.idx)))

	*features = append(*features,
		actName+"+"+prefix+"+dist:"+distStr(dist),
		actName+"+"+prefix+"+dist:"+distStr(dist)+"+leftPos:"+left.posTag+"+rightPos:"+right.posTag,
	)
}

func extractFeatures(state *State, actName string, idx int) []string {
	features := make([]string, 0)
	AddUnigramFeatures(&features, state, actName, idx)
	AddStructuralSingleFeatures(&features, state, actName, idx)

	p0 := NilSafePendingWord(state, idx-1)
	p1 := NilSafePendingWord(state, idx)
	p2 := NilSafePendingWord(state, idx+1)
	p3 := NilSafePendingWord(state, idx+2)

	AddBigramFeatures(&features, actName, p1, p2, "p_i+p_{i+1}")
	AddBigramFeatures(&features, actName, p1, p3, "p_i+p_{i+2}")
	AddBigramFeatures(&features, actName, p0, p1, "p_{i-1}+p_i")
	AddBigramFeatures(&features, actName, p0, p3, "p_{i-1}+p_{i+2}")
	AddBigramFeatures(&features, actName, p2, p3, "p_{i+1}+p_{i+2}")

	addStructuralPairFeatures(&features, actName, p1, p2, "p_i+p_{i+1}")
	addStructuralPairFeatures(&features, actName, p1, p3, "p_i+p_{i+2}")
	addStructuralPairFeatures(&features, actName, p0, p1, "p_{i-1}+p_i")
	addStructuralPairFeatures(&features, actName, p0, p3, "p_{i-1}+p_{i+2}")
	addStructuralPairFeatures(&features, actName, p2, p3, "p_{i+1}+p_{i+2}")

	return features
}

func mod(n, m int) int {
	if n < 0 {
		return (m - (-n % m)) % m
	} else {
		return n % m
	}
}

var MaxFeatureLength = 1000000

func JenkinsHash(s string) int {
	hash := 0
	for _, b := range []byte(s) {
		hash += int(b)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return mod(hash, MaxFeatureLength)
}

func ExtractFeatures(state *State, pair ActionIndexPair) ([]int, error) {
	actName := runtime.FuncForPC(reflect.ValueOf(pair.action).Pointer()).Name()

	featStrs := extractFeatures(state, actName, pair.index)
	features := make([]int, len(featStrs))
	for idx, s := range featStrs {
		features[idx] = JenkinsHash(s)
	}

	return features, nil
}
