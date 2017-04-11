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

	model := Model{make(map[string]float64), make(map[string]float64), 1}

	for iter := 0; iter < 10; iter++ {
		shuffle(sentences)
		goldHeads := make([][]int, 0)
		for _, sent := range sentences {
			goldHeads = append(goldHeads, sent.ExtractHeads())
		}

		for _, sent := range sentences {
			model.Update(sent)
		}

		predHeads := make([][]int, 0)
		w := model.AveragedWeight()
		for _, sent := range sentences {
			Decode(&w, sent)
			predHeads = append(predHeads, sent.ExtractPredictedHeads())
		}
		accuracy, _ := DependencyAccuracy(goldHeads, predHeads)
		fmt.Println(accuracy)
	}
}
