package sorting

// Mergesort - (NlogN running time)
func Mergesort(a *[]int) {
	aux := make([]int, len(*a))
	sort(a, &aux, 0, len(*a)-1)
}

func sort(a *[]int, aux *[]int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid-1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a *[]int, aux *[]int, lo int, mid int, hi int) {
	i, j := lo, mid+1
	// Copy all values from a[lo...hi] into aux[lo...hi]
	for k := lo; k <= hi; i++ {
		(*aux)[k] = (*a)[k]
	}
	// Merge back to a[lo...hi]
	for k := lo; k <= hi; k++ {
		if i > mid {
			j++
			(*a)[k] = (*aux)[j]
		} else if j > hi {
			i++
			(*a)[k] = (*aux)[i]
		} else if (*aux)[j] < (*aux)[i] {
			j++
			(*a)[k] = (*aux)[j]
		} else {
			i++
			(*a)[k] = (*aux)[i]
		}
	}
}
