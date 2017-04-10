package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
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
	weight := make(map[string]float64)

	s := NewState(sent.words)
	decode(&weight, s)
	if len(s.arcs) == 0 {
		t.Error("length of arcs must be greater than 0")
	}
}
