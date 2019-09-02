package main

import "fmt"

// This file contains a collection of useful utility functions

// PrintSlice prints a slice to the console
func PrintSlice(s []int, label string) {
	fmt.Printf("%-8v[", label)
	for _, e := range s {
		fmt.Printf("%3v", e)
	}
	fmt.Printf("%3v", "]\n")
}
