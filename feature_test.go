package main

import (
	"testing"
)

func TestAddUnigramFeatures(t *testing.T) {
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
	features := make([]string, 0)
	AddUnigramFeatures(&features, s, "left", 1)

	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}
}
