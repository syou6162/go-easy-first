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
	s := NewState(words)
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
	s := NewState(words)
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
	s := NewState(words)
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
	s := NewState(words)
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
	s := NewState(words)

	if 10 != len(CandidateActions(s)) {
		t.Error("length of candidate actions must be 10")
	}
}

func TestUpdateWeight(t *testing.T) {
	model := Model{make(map[string]float64), make(map[string]float64), 1}
	gold := []string{"1", "2", "3"}
	predict := []string{"1", "3", "4"}
	model.updateWeight(&gold, &predict)

	if w, _ := model.weight["1"]; w != 0 {
		t.Error("weight of '1' must be 0")
	}
	if w, _ := model.weight["2"]; w != 1 {
		t.Error("weight of '2' must be 1")
	}
	if w, _ := model.weight["3"]; w != 0 {
		t.Error("weight of '3' must be 0")
	}
	if w, _ := model.weight["4"]; w != -1 {
		t.Error("weight of '4' must be -1")
	}

	model.updateWeight(&gold, &predict)

	if w, _ := model.cumWeight["1"]; w != 0 {
		t.Error("cumWeight of '1' must be 0")
	}
	if w, _ := model.cumWeight["2"]; w != 3 {
		t.Error("cumWeight of '2' must be 3")
	}
	if w, _ := model.cumWeight["3"]; w != 0 {
		t.Error("cumWeight of '3' must be 0")
	}
	if w, _ := model.cumWeight["4"]; w != -3 {
		t.Error("cumWeight of '4' must be -3")
	}
}

func TestUpdate(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 1, 2),
		makeWord("hang", "NNP", 2, 3),
		makeWord("plays", "VBZ", 3, 0),
		makeWord("elianti", "NNP", 4, 3),
		makeWord(".", ".", 5, 3),
	)
	sent := Sentence{words: words}
	model := Model{make(map[string]float64), make(map[string]float64), 1}
	model.Update(&sent)
	if model.count == 1 {
		t.Error("count must be greater than 1")
	}
}
