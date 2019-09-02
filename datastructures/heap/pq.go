// Package heap ...
package heap

// TODOs:
// - min PQ (currently a max PQ)
// - make keys should be interfaces

// PriorityQueue ...
type PriorityQueue struct {
	Max, Size int
	Keys      []int
}

// NewPQ constructor
func NewPQ(max int) *PriorityQueue {
	return &PriorityQueue{
		Keys: make([]int, max+1),
		Max:  max,
		Size: 0,
	}
}

// Insert inserts an element into the priority queue
func (pq *PriorityQueue) Insert(key int) {
	pq.Keys[pq.Size] = key
	pq.Size++
	pq.heapifyUp(pq.Size)
}

// ExtractMax removes the max element form the PQ and returns it
func (pq *PriorityQueue) ExtractMax() int {
	max := pq.Keys[1]
	pq.Size--
	pq.exch(1, pq.Size)
	pq.Keys[pq.Size+1] = -1 // TODO: after implementing pointers, this should be set to nil
	pq.heapifyDown(1)
	return max
}

// FindMax Returns the max element (without removal)
func (pq *PriorityQueue) FindMax() int {
	return pq.Keys[1]
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Size == 0
}

// Private APIs

func (pq *PriorityQueue) heapifyDown(k int) {
	for k*2 <= pq.Max {
		j := k * 2
		if j < pq.Max && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *PriorityQueue) heapifyUp(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}
}

func (pq *PriorityQueue) exch(i int, j int) {
	pq.Keys[i], pq.Keys[j] = pq.Keys[j], pq.Keys[i]
}

func (pq *PriorityQueue) less(i int, j int) bool {
	return pq.Keys[i] < pq.Keys[j]
}

// func (pq *PriorityQueue)parent()
