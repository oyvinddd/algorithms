// Package shellsort implements the Shellsort algorithm
package shellsort

/*
Sort runs shellsort on the input set. Shellsort is based
on insertion sort, but it allows for exchanges of elements
further apart (insertion sort only exchange adjacent entries).
The algorithm produce partially sorted arrays are that can
eventually be sorted by regular insertion sort. Shellsort is
most useful when the input set is large in size. Generally
speaking, it is a much faster faster algorithm than insertion
sort and selection sort. It uses no extra space, is great for
arbitrary (not necessarily random) input.

Running time: O(n^(5/3))
*/
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
	a[ai], a[bi] = a[bi], a[ai]
}

func less(a int, b int) bool {
	return a < b
}
