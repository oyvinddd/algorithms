// Package binarysearch implements the binary search algorithm
package binarysearch

/*
Search does a binary search for an element in a sorted array. If
the item is found return it, else return -1. Binary search works
by recursively dividing the array in two, making the "active" part
2^k times smaller every time.

Running time: logarithmic, i.e. O(log(n))
*/
func Search(a []int, needle int) int {
	min, max := 0, len(a)-1
	for min <= max {
		mid := (min + max) / 2
		if a[mid] < needle {
			min = mid + 1
		} else if a[mid] > needle {
			max = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// SearchRec is a recursive alternative to the implementation above
func SearchRec(a []int, min int, max int, needle int) int {
	if max < min {
		return -1
	}
	mid := (min + max) / 2
	if a[mid] == needle {
		return mid
	}
	if a[mid] < needle {
		min = mid + 1
	} else if a[mid] > needle {
		max = mid - 1
	}
	return SearchRec(a, min, max, needle)
}
