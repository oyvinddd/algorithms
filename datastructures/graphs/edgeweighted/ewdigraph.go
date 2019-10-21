package graphs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/bag"
)

// EWDigraph struct
type EWDigraph struct {
	v, e int
	adj  []bag.Bag
}

// NewEWDigraph constructor
func NewEWDigraph(v int) *EWDigraph {
	g := &EWDigraph{v: v, e: 0}
	g.adj = make([]bag.Bag, v)
	for i := 0; i < v; i++ {
		g.adj[i] = *bag.NewBag()
	}
	return g
}

// NoOfV returns the number of vertices in the graph
func (ewg EWDigraph) NoOfV() int {
	return ewg.v
}

// NoOfE returns the number of edges in the graph
func (ewg EWDigraph) NoOfE() int {
	return ewg.e
}

// AddEdge add an edge to the graph
func (ewg *EWDigraph) AddEdge(e Edge) {
	v := e.Either()
	ewg.adj[v].Add(e)
	ewg.e++
}

// GetAdj returns an iterator over the given bag
func (ewg EWDigraph) GetAdj(v int) <-chan interface{} {
	return ewg.adj[v].Iterator()
}

// GetEdges in the graph
func (ewg EWDigraph) GetEdges() {

	for i := 0; i < ewg.NoOfV(); i++ {
		for j := range ewg.GetAdj(i) {
			fmt.Printf("%v", j)
		}
	}
}

// Print graph details to the console
func (ewg EWDigraph) Print() {
	fmt.Printf("%v vertices, %v edges\n", ewg.NoOfV(), ewg.NoOfE())
	for v := 0; v < ewg.NoOfV(); v++ {
		fmt.Printf("%v: ", v)
		for w := range ewg.GetAdj(v) {
			fmt.Printf("%v ", w)
		}
		fmt.Println()
	}
}
