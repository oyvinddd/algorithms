package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/searching/dfs"
	"github.com/oyvinddd/algorithms/datastructures/graphs"
)

func main() {

	// The slice we want to sort
	// a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	g := graphs.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.Print()

	search := dfs.NewDFS()
	search.DepthFirstSearch(g, 0)
	connected := search.IsConnected(g)
	fmt.Printf("Is G connected? Answer: %v\n", connected)
}
