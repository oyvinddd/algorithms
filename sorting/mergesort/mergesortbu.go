package mergesort

/*
MergesortBU performs a bottom-up mergesort. It differs from the recursive
mergesort in that it does all merges on tiny subarrays (1x1, 2x2, ... NxN)
in one pass, then do a second pass to merge those subarrays in pairs, and
so forth, continuing intil we do a merge that encompasses the whole array.

Running time: O(n*log(n))
*/
func MergesortBU(a []int) {
	len := len(a)
	aux = make([]int, len)           // auxiallary array is defined in mergesort.go
	for sz := 1; sz < len; sz *= 2 { // sz: ubarray size
		for lo := 0; lo < len-sz; lo += sz + sz { // lo: subarray index
			merge(a, lo, lo+sz-1, min(lo+sz+sz-1, len-1)) // using merge function from mergesort.go
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
