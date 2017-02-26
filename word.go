package main

type Word struct {
	surface  string
	lemma    string
	posTag   string
	cposTag  string
	idx      int
	head     *int
	children []*Word
}

func makeWord(surface string, posTag string, idx int, head int) *Word {
	return &Word{surface, surface, posTag, posTag, idx, &head, make([]*Word, 0)}
}

func makeRootWord() *Word {
	return makeWord("*ROOT*", "*ROOT*", 0, -1)
}

func (word *Word) appendChild(c *Word) []*Word {
	word.children = append(word.children, c)
	return word.children
}

func (word *Word) prependChild(c *Word) []*Word {
	word.children = append([]*Word{c}, word.children...)
	return word.children
}
