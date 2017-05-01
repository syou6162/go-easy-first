package main

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
)

type FvCache map[string][]string

type State struct {
	pending   []*Word
	arcs      map[int]int
	fvCache   FvCache
	score     float64
	appliedFv []string
}

type States []State

func (slice States) Len() int {
	return len(slice)
}

func (slice States) Less(i, j int) bool {
	return slice[i].score > slice[j].score
}

func (slice States) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (state *State) cacheKeyStr(pair ActionIndexPair) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(pair.action).Pointer()).Name()
	left := state.pending[pair.index]
	right := state.pending[pair.index+1]
	return funcName + ":" + strconv.Itoa(left.idx) + "-" + strconv.Itoa(right.idx)
}

func (state *State) InitFvCache() {
	for _, f := range StateActions {
		for idx := 0; idx < len(state.pending)-1; idx++ {
			pair := ActionIndexPair{f, idx}
			fv, _ := ExtractFeatures(state, pair)
			state.fvCache[state.cacheKeyStr(pair)] = fv
		}
	}
}

func NewState(pending []*Word) *State {
	p := make([]*Word, len(pending))
	copy(p, pending)
	state := State{p, make(map[int]int), FvCache{}, 0.0, make([]string, 0)}
	state.InitFvCache()
	return &state
}

func CopyState(state *State) *State {
	s := &State{append([]*Word{}, state.pending...), make(map[int]int), FvCache{}, 0.0, append([]string{}, state.appliedFv...)}

	for k, v := range state.arcs {
		s.arcs[k] = v
	}

	for k, v := range state.fvCache {
		s.fvCache[k] = v
	}

	return s
}

func (state *State) deletePending(idx int) []*Word {
	state.pending = append(state.pending[:idx], state.pending[idx+1:]...)
	return state.pending
}

func (state *State) ResetFvCache(index int) {
	for _, f := range StateActions {
		min := int(math.Max(0, float64(index-3)))
		max := int(math.Min(float64(len(state.pending)-1), float64(index+3)))
		for idx := min; idx < max; idx++ {
			pair := ActionIndexPair{f, idx}
			delete(state.fvCache, state.cacheKeyStr(pair))
		}
	}
}

func (state *State) GetFvCache(pair ActionIndexPair) []string {
	key := state.cacheKeyStr(pair)
	if fv, ok := state.fvCache[key]; ok {
		return fv
	} else {
		fv, _ = ExtractFeatures(state, pair)
		state.fvCache[key] = fv
		return fv
	}
}
