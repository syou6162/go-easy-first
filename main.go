package main

import (
	"fmt"
	"math/rand"
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
	goldSents, _ := ReadData("/Users/yasuhisa/Desktop/work/easy-first/dev.txt")
	devSents, _ := ReadData("/Users/yasuhisa/Desktop/work/easy-first/test.txt")

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
