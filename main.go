package main

import (
	"fmt"

	"bitbucket.com/oyvind_hauge/inf234/algorithms/sorting/shellsort"
)

func main() {

	// The slice we want to sort
	a := []int{4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	printSlice(a, "INPUT")
	shellsort.Sort(a)
	// insertion.Sort(a)
	printSlice(a, "OUTPUT")
}

func printSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
