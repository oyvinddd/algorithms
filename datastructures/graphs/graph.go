// Package graphs contains an undirected and a directed graph data structure based on an adjacency list
package graphs

import (
	"fmt"
	"log"

	"github.com/oyvinddd/algorithms/datastructures/bag"
)

// Graph struct
type Graph struct {
	v, e int        // Number of vertices and edges
	adj  []*bag.Bag // Adjacency list using bags
}

// NewGraph constructor
func NewGraph(v int) *Graph {
	adj := make([]*bag.Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = bag.NewBag()
	}
	return &Graph{v: v, e: 0, adj: adj}
}

// NewTinyGraph temp constructor until we have gutil up and running
func NewTinyGraph() *Graph {
	g := NewGraph(13)
	g.AddEdge(0, 5)
	g.AddEdge(4, 3)
	g.AddEdge(0, 1)
	g.AddEdge(9, 12)
	g.AddEdge(6, 4)
	g.AddEdge(5, 4)
	g.AddEdge(0, 2)
	g.AddEdge(11, 12)
	g.AddEdge(9, 10)
	g.AddEdge(0, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(5, 3)
	return g
}

// NewTinyGraph2 ...
func NewTinyGraph2() *Graph {
	g := NewGraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 6)
	g.AddEdge(2, 7)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 7)
	g.AddEdge(8, 9)
	g.AddEdge(10, 11)
	g.AddEdge(11, 12)
	return g
}

// NewTinyGraph3 ...
func NewTinyGraph3() *Graph {
	g := NewGraph(6)
	g.AddEdge(0, 5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	return g
}

// AddEdge adds a new edge to the graph
func (g *Graph) AddEdge(v int, w int) {
	g.adj[v].Add(w) // Add w to v's list
	g.adj[w].Add(v) // Add v to w's list
	g.e++
}

// NoOfV returns the number of vertices in the graph
func (g *Graph) NoOfV() int {
	return g.v
}

// NoOfE returns the number of edges in the graph
func (g *Graph) NoOfE() int {
	return g.e
}

// GetAdj returns an iterator over the given bag
func (g *Graph) GetAdj(v int) <-chan interface{} {
	return g.adj[v].Iterator()
}

// Print graph details to the console
func (g *Graph) Print() {
	fmt.Printf("%v vertices, %v edges\n", g.NoOfV(), g.NoOfE())
	for v := 0; v < g.NoOfV(); v++ {
		fmt.Printf("%v: ", v)
		for w := range g.GetAdj(v) {
			fmt.Printf("%v ", w)
		}
		fmt.Println()
	}
}

// Exit with error
func exitWithError(msg string) {
	log.Fatal(msg)
}
