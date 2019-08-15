// Package quicksort implements the Quicksort algorithm
package quicksort

// Sort ....
// divide and conquer, in-place, nlogn
func Sort(a []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	Sort(a, lo, j-1)
	Sort(a, j+1, hi)
}

func partition(a []int, lo int, hi int) int {
	var i, j int
	v := a[lo]
	for true {
		for i = lo + 1; less(a[i], v); i++ {
			if i == hi {
				break
			}
		}
		for j = hi; less(v, a[j]); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		exch(a, i, j)
	}
	exch(a, lo, j)
	return j
}

func less(a int, b int) bool {
	return a < b
}

func exch(a []int, ia int, ib int) {
	a[ia], a[ib] = a[ib], a[ia]
}
