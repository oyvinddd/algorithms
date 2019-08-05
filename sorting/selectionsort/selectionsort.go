package selectionsort

// Sort runs selection sort on the given input slice
// Running time: O(n^2), i.e. quadratic, for all input
func Sort(a []int) {
	len := len(a)
	for i := 0; i < len; i++ {
		minIdx := i
		for j := i + 1; j < len; j++ {
			if less(a[j], a[minIdx]) {
				minIdx = j
			}
		}
		exch(a, i, minIdx)
	}
}

func exch(a []int, indexA int, indexB int) {
	a[indexA], a[indexB] = a[indexB], a[indexA]
}

func less(a int, b int) bool {
	return a < b
}
