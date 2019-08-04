package sorting

// SelectionSort ...
type SelectionSort struct {
	A []int
}

// Execute ...
func (ss *SelectionSort) Execute() {
	ss.sort()
}

func (ss *SelectionSort) sort() {
	arr := ss.A
	len := len(arr)
	for i := 0; i < len; i++ {
		minIdx := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		ss.swap(i, minIdx)
	}
}

func (ss *SelectionSort) swap(indexA int, indexB int) {
	arr := ss.A
	temp := arr[indexA]
	arr[indexA] = arr[indexB]
	arr[indexB] = temp
}
