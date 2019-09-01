package bfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
	"github.com/oyvinddd/algorithms/datastructures/queue"
)

// TODO:
// implement 'check if digraph is strongly connected'
//		pp. 98 in algorithm design (do BFS from s in G and from s in GRev. All nodes should be reachable)

// BFS ...
func BFS(g *graphs.Graph, s int) {
	marked := make([]bool, g.NoOfV())
	edgeTo := make([]int, g.NoOfV())
	queue := queue.NewQueue()
	queue.Enqueue(s)
	marked[s] = true
	for !queue.IsEmpty() {
		v := (*queue.Dequeue()).(int)
		for w := range g.GetAdj(v) {
			ww := w.(int)
			if !marked[ww] {
				fmt.Printf("VISITING EDGE (%v, %v)\n", v, ww)
				edgeTo[ww] = v
				marked[ww] = true
				queue.Enqueue(ww)
			}
		}
	}
}
