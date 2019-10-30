package dp

import (
	"math"

	graphs "github.com/oyvinddd/algorithms/datastructures/graphs/edgeweighted"
)

// BellmanFord runs the Bellman-Ford algorithm on a directed edge-weighted graph
// Running time: O(nm) (for complete graphs the running time is O(n^3))
func BellmanFord(graph *graphs.EWDigraph, source int) {
	V, E := graph.NoOfV(), graph.NoOfE()
	dist := make([]int, V)
	// 1. Initially, set all distances from source to infinity
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt32
	}
	// 2. Relax all edges - |V|-1 times
	for i := 1; i < V-1; i++ {
		for j := 0; j < E; j++ {
			//u := graph[j].
		}
	}
}
