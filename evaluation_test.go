package main

import (
	"testing"
)

func TestDependencyAccuracy(t *testing.T) {
	g1 := []int{1, 2, 3}
	g2 := []int{1, 2, 3}
	g3 := []int{1, 2, 3, 4}
	g := [][]int{g1, g2, g3}

	p1 := []int{1, 2, 30}
	p2 := []int{1, 2, 30}
	p3 := []int{1, 2, 3, 40}
	p := [][]int{p1, p2, p3}

	if a, _ := DependencyAccuracy(g, p); a != 0.7 {
		t.Error("dependency accuracy must be 0.7")
	}
}
