package main

import graphs "github.com/oyvinddd/algorithms/datastructures/graphs/edgeweighted"

func main() {

	g := graphs.NewEWGraph(4)
	g.AddEdge(*graphs.NewEdge(0, 1, 2))
	g.AddEdge(*graphs.NewEdge(1, 2, 1))
	g.AddEdge(*graphs.NewEdge(2, 3, 6))
	g.AddEdge(*graphs.NewEdge(3, 0, 9))

	g.Print()

}
