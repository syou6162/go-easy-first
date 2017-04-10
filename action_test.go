package main

import (
	"testing"
)

func TestAttachLeft(t *testing.T) {
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
	AttachLeft(s, 3)
	p, ok := s.arcs[4]
	if !ok || p != 3 {
		t.Error("parent's index must be 3")
	}

	AttachLeft(s, 3)
	p, ok = s.arcs[5]
	if !ok || p != 3 {
		t.Error("parent's index must be 3")
	}

	if len(s.pending) != 4 {
		t.Error("length of pending must be 4")
	}
}

func TestAttachRight(t *testing.T) {
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
	p, ok := s.arcs[3]
	if !ok || p != 4 {
		t.Error("parent's index must be 4")
	}

	AttachRight(s, 3)
	p, ok = s.arcs[4]
	if !ok || p != 5 {
		t.Error("parent's index must be 5")
	}

	if len(s.pending) != 4 {
		t.Error("length of pending must be 4")
	}
}

func TestAttachLeftAll(t *testing.T) {
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
	AttachLeft(s, 0)
	AttachLeft(s, 0)
	AttachLeft(s, 0)
	AttachLeft(s, 0)
	AttachLeft(s, 0)
	if words[1].surface != "ms." {
		t.Error("surface is wrong")
	}
}
