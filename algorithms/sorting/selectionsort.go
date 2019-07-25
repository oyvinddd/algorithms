package sorting

// SelectionSort ...
func SelectionSort(a *[]int) {
	a1 := *a
	len := len(a1)
	for i := 0; i < len; i++ {
		minIdx := i
		for j := i + 1; j < len; j++ {
			if a1[j] < a1[minIdx] {
				minIdx = j
			}
		}
		swap(a1, i, minIdx)
	}
	a = &a1
}

func swap(arr []int, indexA int, indexB int) {
	temp := arr[indexA]
	arr[indexA] = arr[indexB]
	arr[indexB] = temp
}
