package main

import (
	"testing"
)

func TestAddOneHandFeatures(t *testing.T) {
	features := make([]string, 0)
	word := makeWord("plays", "VBZ", 3, 0)
	addOneHandFeatures(&features, word, "right")
	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}
}

func TestExtractFeature(t *testing.T) {
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
	features := extractFeature(s, 2)
	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}
}
