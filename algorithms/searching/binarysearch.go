package searching

// BinarySearch - search for an element in a sorted array (running time: log(N))
func BinarySearch(a []int, needle int) int {
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

// BinarySearchRec - a recursive variant of the binary search algorithm (running time: log(N))
func BinarySearchRec(a []int, min int, max int, needle int) int {
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
	return BinarySearchRec(a, min, max, needle)
}
