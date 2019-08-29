// Package queue implements a basic FIFO queue using linked-lists
package queue

// Queue struct
type Queue struct {
	first *node
	last  *node
	count int
}

// Unexported
type node struct {
	item *interface{}
	next *node
}

// NewQueue constructor
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue pushes an item onto the queue
func (q *Queue) Enqueue(item interface{}) {
	oldLast := q.last
	q.last = &node{
		item: &item,
		next: nil,
	}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.count++
}

// Dequeue gets the first item in the queue and decrements the size
func (q *Queue) Dequeue() *interface{} {
	if q.first != nil {
		item := q.first.item
		q.first = q.first.next
		q.count--
		if q.IsEmpty() {
			q.last = nil
		}
		return item
	}
	return nil
}

// IsEmpty is true if no items are present in the queue
func (q *Queue) IsEmpty() bool {
	return q.count == 0
}

// Size returns the current size of the queue
func (q *Queue) Size() int {
	return q.count
}
