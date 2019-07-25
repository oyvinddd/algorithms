package sorting

// LinearSort ...
func LinearSort(a []int) []int {
	b := make([]int, len(a))
	for i := range a {
		b[a[i]] = a[i]
	}
	return b
}
