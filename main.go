package main

import (
	"fmt"

	dp "github.com/oyvinddd/algorithms/algorithms/dp/fibonacci"
)

// TODOs:
// 0. get ew graph to return a set of edges
// 1. implement union-find
// 2. implement kruskal (using UF)
// 3. implement prim
// 4. implement dijkstra

func main() {

	m := 40
	n := dp.Fib(m)
	fmt.Printf("Fib(%v) = %v\n", m, n)

	q := dp.FibOpt(m)
	fmt.Printf("FibOpt(%v) = %v\n", m, q)
}
