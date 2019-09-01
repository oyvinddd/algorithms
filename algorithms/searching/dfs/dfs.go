package dfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
	"github.com/oyvinddd/algorithms/datastructures/stack"
)

// TODO:
// implement 'find connected components'
// implement 'find paths in graph'
// implement 'check if digraph has a cycle'

// DFS runs a depth-first search on the graph
// (Returns the number of reachable vertecies from s)
// pp. 93 - Algorithm design
func DFS(graph *graphs.Graph, start int) int {
	explored := make([]bool, graph.NoOfV())
	count := 0
	s := stack.NewStack()
	s.Push(start)
	for !s.IsEmpty() {
		u := (*s.Pop()).(int)
		if !explored[u] {
			fmt.Printf("DFS: Exploring %v...\n", u+1)
			count++
			explored[u] = true
			for v := range graph.GetAdj(u) {
				s.Push(v)
			}
		}
	}
	return count
}

// // IsStronglyConnected checks if a given graph is strongly connected
// // (i.e. if for all pairs (s, t), s is reachable from t and t is reachable from s)
// // linear time: c + (n+m) + (n+m) => c + 2n + 2m => O(n+m)
// func IsStronglyConnected(g *graphs.Digraph) bool {
// 	// Get total number of vertices (constant time)
// 	totalNumberOfV := g.NoOfV()
// 	// Do DFS to count number of reachable vertices from s in G O(n+m)
// 	noOfVInG := DFS(g, 0)
// 	// Do DFS to count number of reachable vertices from s in G-Rev O(n+m)
// 	noOfVInGRev := DFS(g.Reverse(), 0)
// 	return (totalNumberOfV == noOfVInG) && (totalNumberOfV == noOfVInGRev) && (noOfVInG == noOfVInGRev)
// }

// // HasCycle uses DFS to check if a given digraph has a cycle
// //
// func HasCycle(g *graphs.Digraph) bool {
// 	recStack := make([]bool, g.NoOfV())
// 	return false
// }
