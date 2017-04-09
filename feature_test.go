package main

import (
	"reflect"
	"testing"
)

func TestAddOneHandFeatures(t *testing.T) {
	features := make([]string, 0)
	word := makeWord("plays", "VBZ", 3, 0)
	addOneHandFeatures(&features, word, RIGHT)

	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}

	rightFeatures := []string{"side:1+surface:plays", "side:1+lemma:plays", "side:1+posTag:VBZ", "side:1+cposTag:VBZ"}
	if !reflect.DeepEqual(features, rightFeatures) {
		t.Error("feature extraction seems wrong")
	}
}

func TestExtractAttachLeftFeature(t *testing.T) {
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
	features := extractAttachLeftFeatures(s, 2)
	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}
}

func TestExtractAttachRightFeature(t *testing.T) {
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
	features := extractAttachRightFeatures(s, 2)
	if len(features) == 0 {
		t.Error("length of features must be greater than 0")
	}
}
