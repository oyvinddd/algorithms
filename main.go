package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/dp"
)

// TODOs:
// 0. get ew graph to return a set of edges
// 1. implement union-find
// 2. implement kruskal (using UF)
// 3. implement prim
// 4. implement dijkstra

func main() {
	lcs := dp.LongestCommonSubstring("ABC", "BABA")
	fmt.Println(lcs)
}
