// Package selectionsort implements the selection sort algorithm
package selectionsort

/*
Sort runs selection sort on the input set. For every iteration
of the outer loop, the inner loop checks every remaining element
to see if there exists a smaller element than the one at position
i. If one such element exists, swap the two elements. Sorting is
done in-place (i.e. no extra space is needed).

Running time: O(n^2), i.e. quadratic, for all input
*/
func Sort(a []int) {
	for i := range a[:len(a)] {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[k] {
				k = j
			}
		}
		a[i], a[k] = a[k], a[i]
	}
}
