package main

import (
	"fmt"

	"bitbucket.com/oyvind_hauge/inf234/algorithms/sorting"
)

func main() {

	a := []int{4, 7, 2, 1, 6, 7, 1, 9, 8}

	fmt.Printf("UNSORTED: %v\n", a)
	(&sorting.SelectionSort{A: a}).Execute()
	fmt.Printf("SORTED: %v\n", a)
}
