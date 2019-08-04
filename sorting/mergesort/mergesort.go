package mergesort

var aux []int

// Mergesort - (NlogN running time)
func Mergesort(a *[]int) {
	aux = make([]int, len(*a))
	sort(a, 0, len(*a)-1)
}

func sort(a *[]int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, lo, mid)
	sort(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func merge(a *[]int, lo int, mid int, hi int) {
	i, j := lo, mid+1
	// Copy all values from a[lo...hi] into aux[lo...hi]
	for k := lo; k <= hi; k++ {
		aux[k] = (*a)[k]
	}
	// Merge back to a[lo...hi]
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*a)[k] = aux[j]
			j++
		} else if j > hi {
			(*a)[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			(*a)[k] = aux[j]
			j++
		} else {
			(*a)[k] = aux[i]
			i++
		}
	}
}
