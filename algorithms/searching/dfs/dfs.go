package dfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
)

var marked []bool
var count int

// DepthFirstSearch ...
func DepthFirstSearch(g *graphs.Graph, s int) {
	marked = make([]bool, g.NoOfV())
	dfs(g, s)
}

func dfs(g *graphs.Graph, v int) {
	fmt.Printf("DFS: visiting node: %v\n", v)
	marked[v] = true
	count++
	for w := range g.GetAdj(v) {
		if !marked[w.(int)] {
			dfs(g, w.(int))
		}
	}
}
