package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
)

func shuffle(data []*Sentence) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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

	splitPos := int(float64(len(sentences)) * 0.8)
	goldSents := sentences[0:splitPos]
	devSents := sentences[splitPos+1 : len(sentences)-1]

	model := NewModel()

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
