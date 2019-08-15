// Package datastructures contains various basic data structures
package datastructures

// Stack struct
type Stack struct {
	first *node
	count int
}

type node struct {
	item *string
	next *node
}

// Push an item onto the stack
func (s *Stack) Push(item string) {
	oldFirst := s.first
	s.first = &node{item: &item}
	s.first.next = oldFirst
	s.count++
}

// Pop an item from the stack
func (s *Stack) Pop() *string {
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
