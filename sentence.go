package main

type Sentence struct {
	words []*Word
}

// extract heads without root for evaluation
func (sent *Sentence) ExtractHeads() []int {
	heads := make([]int, 0)
	for _, w := range sent.words[1:] {
		heads = append(heads, w.head)
	}
	return heads
}

func (sent *Sentence) ExtractPredictedHeads() []int {
	heads := make([]int, 0)
	for _, w := range sent.words[1:] {
		heads = append(heads, w.predHead)
	}
	return heads
}
