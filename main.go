package main

import (
	"fmt"

	"bitbucket.com/oyvind_hauge/inf234/algorithms/searching"
)

func main() {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	index := searching.BinarySearch(primes, 67)
	idx2 := searching.BinarySearchRec(primes, 0, len(primes)-1, 67)

	fmt.Println(index)
	fmt.Println(idx2)
}
