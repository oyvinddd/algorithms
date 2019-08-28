package dfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
	"github.com/oyvinddd/algorithms/datastructures/stack"
)

// TODO:
// implement 'find connected components'
// implement 'find paths in graph'

// DFS runs a depth-first search on the graph
// pp. 93 - Algorithm design
func DFS(graph *graphs.Graph) {
	explored := make([]bool, graph.NoOfV())
	s := stack.NewStack()
	s.Push(0)
	for !s.IsEmpty() {
		u := (*s.Pop()).(int)
		if !explored[u] {
			fmt.Printf("DFS: Exploring %v...\n", u)
			explored[u] = true
			for v := range graph.GetAdj(u) {
				s.Push(v)
			}
		}
	}
}
