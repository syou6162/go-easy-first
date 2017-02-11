package main

type State struct {
	pending []*Word
	arcs    map[int]int
}
