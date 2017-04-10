package main

import (
	"testing"
)

func TestDeletePending(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)
	s := NewState(words)
	s.deletePending(2)

	if s.pending[1].surface != "ms." {
		t.Error("surface must be 'ms.'")
	}
	if s.pending[2].surface != "plays" {
		t.Error("surface must be 'plays'")
	}

	s.deletePending(1)
	if s.pending[1].surface != "plays" {
		t.Error("surface must be 'plays'")
	}

	if words[1].surface != "ms." {
		t.Error("surface is wrong!!!" + words[1].surface)
	}
}
