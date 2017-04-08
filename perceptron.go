package main

import (
	"errors"
	"math"
)

// GoldArcs returns map of parent => children
func GoldArcs(sent *Sentence) map[int][]int {
	result := make(map[int][]int)
	for idx, w := range sent.words {
		head := *w.head
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
	return 0.0
}

func BestActionIndexPair(weight *map[string]float64, state *State) ActionIndexPair {
	bestScore := math.Inf(-1)
	pairs := CandidateActions(state)
	bestPair := pairs[0]
	for pairIdx, pair := range pairs {
		idx := pair.index
		score := 0.0
		if pairIdx%2 == 0 { // AttachLeft
			score = DotProduct(weight, extractAttachLeftFeature(state, idx))
		} else { // AttachRight
			score = DotProduct(weight, extractAttachRightFeature(state, idx))
		}
		if score > bestScore {
			bestPair = pair
		}
	}
	return bestPair
}
