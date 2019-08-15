package main

import (
	"fmt"

	stack "github.com/oyvinddd/algorithms/data-structures/stack"
)

func main() {

	// The slice we want to sort
	a := []int{9, 4, 7, 5, 0, 3, 2, 6, 1, 9, 8}

	printSlice(a, "INPUT")

	// Run code on slice here
	// s := stack.NewStack()
	s := stack.Stack{}
	s.Push("hello")
	s.Push("world")
	s.Push("!!!!!")

	for s.Size() > 0 {
		fmt.Printf("%v (size: %v)\n", *s.Pop(), s.Size())
	}
	ss := s.Pop()
	fmt.Print(ss)

	printSlice(a, "OUTPUT")
}

func printSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
