package main

import (
	"errors"
	"math"
	"reflect"
)

// GoldArcs returns map of parent => children
func GoldArcs(sent *Sentence) map[int][]int {
	result := make(map[int][]int)
	for idx, w := range sent.words {
		head := w.head
		if children, ok := result[head]; ok {
			result[head] = append(children, idx)
		} else {
			result[head] = []int{idx}
		}
	}
	return result
}

// EdgeFor returns a pair of parent index and child index
func EdgeFor(state *State, actionID int, idx int) ([]int, error) {
	switch actionID {
	case 0:
		return []int{state.pending[idx].idx, state.pending[idx+1].idx}, nil
	case 1:
		return []int{state.pending[idx+1].idx, state.pending[idx].idx}, nil
	default:
		return nil, errors.New("Invalid line")
	}
}

// IsValid returns the chosen action/location pair is valid
func IsValid(state *State, actionID int, idx int, goldArcs map[int][]int) bool {
	pair, err := EdgeFor(state, actionID, idx)
	if err != nil {
		return false
	}
	pIdx := pair[0]
	cIdx := pair[1]
	containedInGoldArcs := false
	for _, i := range goldArcs[pIdx] {
		if cIdx == i {
			containedInGoldArcs = true
			break
		}
	}
	flag := false
	for _, cPrime := range goldArcs[cIdx] {
		if cIdx != state.arcs[cPrime] {
			flag = true
			break
		}
	}
	if !containedInGoldArcs || flag {
		return false
	}
	return true
}

type ActionIndexPair struct {
	action StateAction
	index  int
}

func (pair1 ActionIndexPair) SameActionIndexPair(pair2 ActionIndexPair) bool {
	return pair1.index == pair2.index &&
		reflect.ValueOf(pair1.action).Pointer() == reflect.ValueOf(pair2.action).Pointer()
}

func AllowedActions(state *State, goldArcs map[int][]int) []ActionIndexPair {
	result := make([]ActionIndexPair, 0)
	for actionID, f := range StateActions {
		for idx := 0; idx < len(state.pending)-1; idx++ {
			if IsValid(state, actionID, idx, goldArcs) {
				result = append(result, ActionIndexPair{f, idx})
			}
		}
	}
	return result
}

func CandidateActions(state *State) []ActionIndexPair {
	result := make([]ActionIndexPair, 0)
	for _, f := range StateActions {
		for idx := 0; idx < len(state.pending)-1; idx++ {
			result = append(result, ActionIndexPair{f, idx})
		}
	}
	return result
}

func DotProduct(weight *map[string]float64, fv []string) float64 {
	sum := 0.0
	for _, f := range fv {
		if v, ok := (*weight)[f]; ok {
			sum += v
		}
	}
	return sum
}

func BestActionIndexPair(weight *map[string]float64, state *State) ActionIndexPair {
	bestScore := math.Inf(-1)
	pairs := CandidateActions(state)
	bestPair := pairs[0]
	for _, pair := range pairs {
		score := 0.0
		fv, _ := ExtractFeatures(state, pair)
		score = DotProduct(weight, fv)
		if score > bestScore {
			bestPair = pair
			bestScore = score
		}
	}
	return bestPair
}

func BestAllowedActionIndexPair(weight *map[string]float64, state *State, pairs []ActionIndexPair) ActionIndexPair {
	bestScore := math.Inf(-1)
	bestPair := pairs[0]
	for _, pair := range pairs {
		fv, _ := ExtractFeatures(state, pair)
		score := DotProduct(weight, fv)
		if score > bestScore {
			bestPair = pair
			bestScore = score
		}
	}
	return bestPair
}

type Model struct {
	weight    map[string]float64
	cumWeight map[string]float64
	count     int
}

func (model *Model) updateWeight(goldFeatureVector *[]string, predictFeatureVector *[]string) {
	for _, feat := range *goldFeatureVector {
		w, _ := model.weight[feat]
		cumW, _ := model.cumWeight[feat]
		model.weight[feat] = w + 1.0
		model.cumWeight[feat] = cumW + float64(model.count)
	}
	for _, feat := range *predictFeatureVector {
		w, _ := model.weight[feat]
		cumW, _ := model.cumWeight[feat]
		model.weight[feat] = w - 1.0
		model.cumWeight[feat] = cumW - float64(model.count)
	}
	model.count += 1
}

func (model *Model) Update(gold *Sentence) {
	state := &State{gold.words, make(map[int]int)}
	goldArcs := GoldArcs(gold)
	iter := 0
	for {
		if len(state.pending) <= 1 {
			break
		}
		allow := AllowedActions(state, goldArcs)
		choice := BestActionIndexPair(&model.weight, state)
		containChoice := false
		//fmt.Println(choice)
		for _, pair := range allow {
			if pair.SameActionIndexPair(choice) {
				containChoice = true
			}
		}
		if containChoice {
			choice.action(state, choice.index)
		} else {
			predFv, _ := ExtractFeatures(state, choice)
			good := BestAllowedActionIndexPair(&model.weight, state, allow)
			goodFv, _ := ExtractFeatures(state, good)
			model.updateWeight(&goodFv, &predFv)
		}
		iter++
		if iter > 500 { // for infinite loop
			break
		}
	}
}
