// Package mergesort implements the merge sort algorithm
package mergesort

var aux []int

/*
Mergesort is a recursive algorithm that divides the input set
into two equal-sized parts, sorts each half recursively, and
then merges the two parts into a single sorted output list.

Running time: O(n log n), i.e. logarithmic
*/
func Mergesort(a []int) {
	len := len(a)
	aux = make([]int, len)
	sort(a, 0, len-1)
}

func sort(a []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, lo, mid)
	sort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func merge(a []int, lo int, mid int, hi int) {
	i, j := lo, mid+1
	// Copy all values from a[lo...hi] into aux[lo...hi]
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	// Merge back to a[lo...hi]
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
