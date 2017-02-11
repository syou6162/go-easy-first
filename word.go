package main

type Word struct {
	surface string
	lemma   string
	posTag  string
	cposTag string
	idx     int
	head    int
}

func makeWord(surface string, posTag string, idx int, head int) *Word {
	return &Word{surface, surface, posTag, posTag, idx, head}
}
