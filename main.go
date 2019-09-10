package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/greedy"
)

func main() {

	j1 := *greedy.NewJob("J1", 0, 2)
	j2 := *greedy.NewJob("J2", 3, 4)
	j3 := *greedy.NewJob("J3", 5, 6)
	j4 := *greedy.NewJob("J4", 5, 7)
	j5 := *greedy.NewJob("J5", 6, 11)

	jobs := []greedy.Job{j1, j2, j3, j4, j5}

	optimalJobs := greedy.IntervalScheduling(jobs)

	for _, e := range optimalJobs {
		fmt.Printf("%v\n", e.Name)
	}
}
