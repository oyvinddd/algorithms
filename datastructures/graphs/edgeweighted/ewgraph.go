package graphs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/bag"
)

// EWGraph struct
type EWGraph struct {
	v, e int
	adj  []bag.Bag
}

// NewEWGraph constructor
func NewEWGraph(v int) *EWGraph {
	g := &EWGraph{v: v, e: 0}
	g.adj = make([]bag.Bag, v)
	for i := 0; i < v; i++ {
		g.adj[i] = *bag.NewBag()
	}
	return g
}

// NoOfV returns the number of vertices in the graph
func (ewg EWGraph) NoOfV() int {
	return ewg.v
}

// NoOfE returns the number of edges in the graph
func (ewg EWGraph) NoOfE() int {
	return ewg.e
}

// AddEdge add an edge to the graph
func (ewg *EWGraph) AddEdge(e Edge) {
	v := e.Either()
	w := e.Other(v)
	ewg.adj[v].Add(e)
	ewg.adj[w].Add(e)
	ewg.e++
}

// GetAdj returns an iterator over the given bag
func (ewg EWGraph) GetAdj(v int) <-chan interface{} {
	return ewg.adj[v].Iterator()
}

// Print graph details to the console
func (ewg EWGraph) Print() {
	fmt.Printf("%v vertices, %v edges\n", ewg.NoOfV(), ewg.NoOfE())
	for v := 0; v < ewg.NoOfV(); v++ {
		fmt.Printf("%v: ", v)
		for w := range ewg.GetAdj(v) {
			fmt.Printf("%v ", w)
		}
		fmt.Println()
	}
}
