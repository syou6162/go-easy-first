package main

import (
	"reflect"
	"testing"
)

func TestEdgeFor(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)
	s := &State{words, make(map[int]int)}
	pair, err := EdgeFor(s, 0, 0)
	if err != nil {
		t.Error("error should be nil")
	}
	if !reflect.DeepEqual(pair, []int{0, 1}) {
		t.Error("pair shoud be [0, 1] but: ", pair)
	}
}

func TestIsValidFalse(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)

	s := &State{words, make(map[int]int)}
	goldArcs := make(map[int][]int)
	goldArcs[-1] = []int{0}
	goldArcs[0] = []int{1}
	goldArcs[1] = []int{2}
	if IsValid(s, 0, 0, goldArcs) != false {
		t.Error("should return false")
	}
}

func TestIsValidTrue(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)

	arcs := make(map[int]int)
	arcs[2] = 1
	s := &State{words, arcs}
	goldArcs := make(map[int][]int)
	goldArcs[-1] = []int{0}
	goldArcs[0] = []int{1}
	goldArcs[1] = []int{2}
	if IsValid(s, 0, 0, goldArcs) != true {
		t.Error("should return true")
	}
}

func TestAllowedActions(t *testing.T) {
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
	AttachRight(s, 3)

	goldArcs := make(map[int][]int)
	goldArcs[-1] = []int{0}
	goldArcs[0] = []int{1}
	goldArcs[1] = []int{2}

	if 1 != len(AllowedActions(s, goldArcs)) {
		t.Error("length of allowed actions must be 1")
	}
}

func TestCandidateActions(t *testing.T) {
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

	if 10 != len(CandidateActions(s)) {
		t.Error("length of candidate actions must be 10")
	}
}
