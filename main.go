package main

import (
	"fmt"

	"github.com/oyvinddd/algorithms/algorithms/other/galeshapeley"
)

func main() {

	// The slice we want to sort
	a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	printSlice(a, "INPUT")

	// Run code here
	m, w := galeshapeley.GetPeople()
	galeshapeley.GaleShapeley(m, w)

	printSlice(a, "OUTPUT")
}

func printSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
