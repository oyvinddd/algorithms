package greedy

import (
	"fmt"

	graphs "github.com/oyvinddd/algorithms/datastructures/graphs/edgeweighted"
)

// TODO: implement this

// KruskalMST ...
func KruskalMST(ewg graphs.EWGraph) []graphs.Edge {
	edges := EdgeSort(ewg)
	mst := make([]graphs.Edge, len(edges))
	for edge := range edges {
		fmt.Printf("%v", edge)
	}
	return mst
}

// EdgeSort sorts edges in the graph from lightest to heaviest weight
// To keep it simple, consider the running time of this to be O(nlogn)
func EdgeSort(ewg graphs.EWGraph) []graphs.Edge {
	// sort.Slice()
	return make([]graphs.Edge, 0)
}
