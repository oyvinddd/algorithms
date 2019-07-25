package sorting

// SelectionSort ...
func SelectionSort(a []int) {
	len := len(a)
	for i := 0; i < len; i++ {
		minIdx := i
		for j := i + 1; j < len; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		swap(a, i, minIdx)
	}
}

func swap(arr []int, indexA int, indexB int) {
	temp := arr[indexA]
	arr[indexA] = arr[indexB]
	arr[indexB] = temp
}
