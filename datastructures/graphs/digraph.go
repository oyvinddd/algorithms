// Package graphs contains an undirected and a directed graph data structure based on an adjacency list
package graphs

import (
	"github.com/oyvinddd/algorithms/datastructures/bag"
)

// Digraph struct
type Digraph struct {
	v, e int
	adj  []*bag.Bag
}

// NewDigraph constructs a new directed graph
func NewDigraph(v int) *Digraph {
	adj := make([]*bag.Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = bag.NewBag()
	}
	return &Digraph{adj: adj}
}

// AddEdge adds a new edge to the graph
func (g *Digraph) AddEdge(v int, w int) {
	g.adj[v].Add(w)
	g.e++
}

// NoOfV returns the number of vertices in the graph
func (g *Digraph) NoOfV() int {
	return g.v
}

// NoOfE returns the number of edges in the graph
func (g *Digraph) NoOfE() int {
	return g.e
}

// GetAdj returns an iterator over the given bag
func (g *Digraph) GetAdj(v int) <-chan interface{} {
	return g.adj[v].Iterator()
}

// Reverse returns a reversed digraph
func (g *Digraph) Reverse() *Digraph {
	newDigraph := NewDigraph(g.v)
	for v := 0; v < g.v; v++ {
		for w := range g.GetAdj(v) {
			newDigraph.AddEdge(w.(int), v)
		}
	}
	return newDigraph
}
