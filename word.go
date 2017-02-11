package main

type Word struct {
	surface  string
	lemma    string
	posTag   string
	cposTag  string
	idx      int
	head     int
	children []*Word
}

func makeWord(surface string, posTag string, idx int, head int) *Word {
	return &Word{surface, surface, posTag, posTag, idx, head, make([]*Word, 0)}
}

func makeRootWord() *Word {
	return makeWord("*ROOT*", "*ROOT*", 0, -1)
}
