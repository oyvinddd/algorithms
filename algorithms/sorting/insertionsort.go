package sorting

// InsertionSort - Elementary sort. Running time of insetion sort relies on the input. If it's already partially sorted, it's faster to sort than using selection sort.
// just like selection sort, this algorithm is quadratic.
func InsertionSort(a []int) {
	len := len(a)
	for i := 1; i < len; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			exchange(a, j, j-1)
		}
	}
}

func exchange(a []int, ai int, bi int) {
	temp := a[ai]
	a[ai] = a[bi]
	a[bi] = temp
}
