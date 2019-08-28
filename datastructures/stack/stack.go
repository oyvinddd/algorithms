// Package stack implements a pushdown LIFO stack using linked-lists
package stack

// Stack struct
type Stack struct {
	first *node
	count int
}

type node struct {
	item *interface{}
	next *node
}

// NewStack constructor
func NewStack() *Stack {
	return &Stack{}
}

// Push an item onto the stack
func (s *Stack) Push(item interface{}) {
	oldFirst := s.first
	s.first = &node{item: &item}
	s.first.next = oldFirst
	s.count++
}

// Pop an item from the stack
func (s *Stack) Pop() *interface{} {
	if s.first != nil {
		item := s.first.item
		s.first = s.first.next
		s.count--
		return item
	}
	return nil
}

// IsEmpty checks if the stack has any items
func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

// Size returns the current size of the stack
func (s *Stack) Size() int {
	return s.count
}
