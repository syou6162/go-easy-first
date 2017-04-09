package main

import (
	"reflect"
	"testing"
)

func TestExtractHeads(t *testing.T) {
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
	head := sent.ExtractHeads()

	if !reflect.DeepEqual(head, []int{2, 3, 0, 3, 3}) {
		t.Error("head extraction seems wrong")
	}
}
