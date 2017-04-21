package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func shuffle(data []*Sentence) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

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

	goldSents := make([]*Sentence, 0)
	devSents := make([]*Sentence, 0)
	for i := 0; i < len(sentences) - 1; i++ {
		if i < int(float64(len(sentences)) * 0.8) {
			goldSents = append(goldSents, sentences[i])
		} else {
			devSents = append(devSents, sentences[i])
		}
	}

	model := Model{make(map[string]float64), make(map[string]float64), 1}

	for iter := 0; iter < 10; iter++ {
		shuffle(goldSents)
		for _, sent := range goldSents {
			model.Update(sent)
		}

		trainAccuracy := DependencyAccuracy(&model, goldSents)
		devAccuracy := DependencyAccuracy(&model, devSents)
		fmt.Println(fmt.Sprintf("%d, %0.03f, %0.03f", iter, trainAccuracy, devAccuracy))
	}
}
