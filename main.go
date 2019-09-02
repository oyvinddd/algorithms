package main

import (
	"github.com/oyvinddd/algorithms/algorithms/searching/dfs"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
)

func main() {

	// The slice we want to sort
	// a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	g := graphs.NewGraph(2)
	g.AddEdge(0, 1)
	g.Print()

	dfs.DepthFirstSearch(g, 0)
}
