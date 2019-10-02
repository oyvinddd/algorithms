package main

import graphs "github.com/oyvinddd/algorithms/datastructures/graphs/edgeweighted"

func main() {

	g := graphs.NewEWGraph(4)
	g.AddEdge(*graphs.NewEdge(0, 1, 2))
	g.AddEdge(*graphs.NewEdge(1, 2, 1))
	g.AddEdge(*graphs.NewEdge(2, 3, 6))
	g.AddEdge(*graphs.NewEdge(3, 0, 9))

	// g.Print()
	g.GetEdges()

	// TODOs:
	// 0. get ew graph to return a set of edges
	// 1. implement union-find
	// 2. implement kruskal (using UF)
	// 3. implement prim
	// 4. implement dijkstra
}
