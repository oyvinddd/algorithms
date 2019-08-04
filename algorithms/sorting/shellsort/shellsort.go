package shellsort

// Sort - ShellSort
// based on insertion sort, allows exchanges of elements far apart (insertion sort only exch adjacent entries)
// produce partially sorted array that can evcentually be sorted by regular insertion sort
// this algo is most useful when the input array is large in size
// MUCH faster than insertion sort and selection sort
// uses no extra space, is great for arbitrary input (not necessarily random)
func Sort(a []int) {
	len := len(a)
	h := 1
	for h < (len / 3) {
		h = h*3 + 1 // 1, 4, 13, 40, 121, ...
	}
	for h >= 1 {
		// h-sort the array
		for i := h; i < len; i++ {
			// Insert a[i] among a[i-h], a[i-h*2], a[i-h*3], etc.
			for j := i; j >= h && less(a[j], a[j-h]); j -= h {
				exch(a, j, j-h)
			}
		}
		h = h / 3
	}
}

func exch(a []int, ai int, bi int) {
	temp := a[ai]
	a[ai] = a[bi]
	a[bi] = temp
}

func less(a int, b int) bool {
	return a < b
}
