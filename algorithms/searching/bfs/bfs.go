package bfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
	"github.com/oyvinddd/algorithms/datastructures/queue"
)

// BFS ...
func BFS(graph *graphs.Graph) {
	discovered := make([]bool, graph.NoOfV())
	bfsTree := make([]*queue.Queue, graph.NoOfV())
	layers := make([]*queue.Queue, graph.NoOfV())
	layer := 0
	for i := 0; i < graph.NoOfV(); i++ {
		discovered[i] = false
		bfsTree[i] = queue.NewQueue()
		layers[i] = queue.NewQueue()
	}
	discovered[0] = true
	layers[0].Enqueue(0)
	for !layers[layer].IsEmpty() {
		u := (*layers[layer].Dequeue()).(int)
		fmt.Printf("U: %v\n", u+1)
		for v := range graph.GetAdj(u) {
			fmt.Printf("V: %v\n", v.(int)+1)
			if !discovered[v.(int)] {
				fmt.Printf("BFS: Exploring %v...\n", v.(int)+1)
				discovered[v.(int)] = true
				bfsTree[u].Enqueue(v.(int))
				layers[layer+1].Enqueue(v.(int))
			}
		}
		layer++
	}
}
