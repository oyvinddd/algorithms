package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/searching/dfs"
	graphs "github.com/oyvinddd/algorithms/datastructures/graphs"
)

func main() {

	// The slice we want to sort
	// a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}
	g := graphs.NewTinyGraph()
	// gutil.PrintGraph(g)
	dfs.DFS(g)
}

func printSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
