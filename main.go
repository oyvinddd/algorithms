package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/greedy"
)

// TODOs:
// 0. get ew graph to return a set of edges
// 1. implement union-find
// 2. implement kruskal (using UF)
// 3. implement prim
// 4. implement dijkstra

func main() {

	j1 := *greedy.NewJob("Job #1", 1, 5)
	j2 := *greedy.NewJob("Job #2", 2, 4)
	j3 := *greedy.NewJob("Job #3", 4, 6)
	j4 := *greedy.NewJob("Job #4", 6, 7)

	jobs := []greedy.Job{j1, j2, j3, j4}

	comp := greedy.IntervalScheduling(jobs)
	fmt.Printf("%v", comp)
}
