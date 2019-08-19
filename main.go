package main

import (
	"fmt"

	b "github.com/oyvinddd/algorithms/datastructures/bag"
)

func main() {

	// The slice we want to sort
	a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	printSlice(a, "INPUT")

	// Run code here
	bag := b.NewBag()
	bag.Add(10)
	bag.Add(2)

	it := bag.BagIterator()
	for item := range it {
		fmt.Printf("%v\n", *item)
	}

	printSlice(a, "OUTPUT")
}

func printSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
