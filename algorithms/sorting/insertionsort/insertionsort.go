package insertionsort

/*
Sort sorts the input slice in-place
Running time: O(n^2), i.e. quadratic, but performs much better than selection sort in most situations
*/
func Sort(a []int) {
	len := len(a)
	for i := 1; i < len; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			exch(a, j, j-1)
		}
	}
}

func exch(a []int, indexA int, indexB int) {
	a[indexA], a[indexB] = a[indexB], a[indexA]
}
