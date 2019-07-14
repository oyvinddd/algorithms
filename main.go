package main

import (
	"fmt"

	sorting "bitbucket.com/oyvind_hauge/inf234/algorithms/sorting"
)

func main() {

	a := []int{4, 7, 2, 1, 6, 7, 1, 9, 8}
	sorting.Mergesort(&a)
	fmt.Printf("%v", a)
}
