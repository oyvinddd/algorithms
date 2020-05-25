package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/dac"
)

// TODOs:
// 0. get ew graph to return a set of edges
// 1. implement union-find
// 2. implement kruskal (using UF)
// 3. implement prim
// 4. implement dijkstra

func main() {
	a := []int{1, 4, 6}
	b := []int{2, 3, 7}
	m, c := dac.MergeAndCount(a, b)
	fmt.Println(m)
	fmt.Println(c)
}
