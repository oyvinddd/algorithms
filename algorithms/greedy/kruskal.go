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
func EdgeSort(ewg graphs.EWGraph) []graphs.Edge {
	return make([]graphs.Edge, 0)
}
