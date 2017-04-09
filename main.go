package main

import (
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	sentences := make([]*Sentence, 0)
	for _, sent := range splitBySentence(string(data)) {
		s, err := makeSentence(sent)
		if err != nil {
			break
		}
		sentences = append(sentences, s)
	}

	model := Model{make(map[string]float64), make(map[string]float64), 1}
	for _, sent := range sentences {
		model.Update(sent)
	}

	for _, sent := range sentences {
		Decode(&model.weight, sent)
	}
}
