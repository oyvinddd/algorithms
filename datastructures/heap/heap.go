package heap

// heap = balanced binary tree

// Heap struct
type Heap struct {
	size   int
	values []int
}

// NewHeap constructor
func NewHeap(size int) *Heap {
	v := make([]int, size)
	return &Heap{size: size, values: v}
}

// public API

func (heap *Heap) Insert(i int) {
}

func (heap *Heap) FindMin() int {
	return -1
}

func (heap *Heap) ExtractMin() int {
	return -1
	// TODO:
}

func (heap *Heap) Delete() {
	// TODO:
}

// private API

func (heap *Heap) heapifyUp(i int) { // O(logn)
	if i > 1 {
		j := heap.parent(i)
		// pp. 61
	}
}

func (heap *Heap) heapifyDown(i int) { // O(logn)

}

func (heap *Heap) parent(i int) int {
	return heap.values[i/2]
}
