package main

import (
	"errors"
)

func dependencyAccuracy(golds [][]int, predictions [][]int) (float64, error) {
	if len(golds) != len(predictions) {
		return 0.0, errors.New("length of golds and that of predictions is not same")
	}
	sum := 0.0
	count := 0.0
	for idx, gold := range golds {
		pred := predictions[idx]
		if len(gold) != len(pred) {
			return 0.0, errors.New("length of gold and that of pred is not same")
		}
		for i, g := range gold {
			if g == pred[i] {
				sum += 1.0
			}
			count += 1.0
		}
	}
	return sum / count, nil
}

func DependencyAccuracy(model *Model, sents []*Sentence) float64 {
	goldHeads := make([][]int, 0)
	for _, sent := range sents {
		goldHeads = append(goldHeads, sent.ExtractHeads())
	}

	predHeads := make([][]int, 0)
	w := model.AveragedWeight()
	for _, sent := range sents {
		Decode(&w, sent)
		predHeads = append(predHeads, sent.ExtractPredictedHeads())
	}
	accuracy, _ := dependencyAccuracy(goldHeads, predHeads)
	return accuracy
}
