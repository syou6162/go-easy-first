package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	goldHeads := make([][]int, 0)
	sentences := make([]*Sentence, 0)
	for _, sent := range splitBySentence(string(data)) {
		s, err := makeSentence(sent)
		if err != nil {
			break
		}
		sentences = append(sentences, s)
	}

	for _, sent := range sentences {
		goldHeads = append(goldHeads, sent.ExtractHeads())
	}

	model := Model{make(map[string]float64), make(map[string]float64), 1}

	for iter := 0; iter < 10; iter++ {
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
