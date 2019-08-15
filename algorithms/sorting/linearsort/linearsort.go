// Package linearsort implements the linear sort algorithm
package linearsort

/*
Sort returns a sorted copy of the input slice. This algorithm
is pretty much pointless and only works for unique, positive
integers since it uses these values as indices when sorting.

Running time: linear time proprtional to the size of the input
*/
func Sort(a []int) []int {
	b := make([]int, len(a))
	for index := range a {
		b[a[index]] = a[index]
	}
	return b
}
