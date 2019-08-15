// Package stack implements a basic (resizing) LIFO stack
package stack

// Stack a custom stack type that stores string pointers
type Stack []*string

// NewStack constructor
func NewStack() *Stack {
	return &Stack{}
}

// Push pushes a new element onto the stack
func (s *Stack) Push(e string) {
	(*s) = append(*s, &e)
}

// Pop reduces the size of the stack by one and returns the last element
func (s *Stack) Pop() *string {
	stack := *s
	if len(stack) > 0 {
		l := len(stack) - 1
		e := stack[l]
		*s = stack[:l]
		return e
	}
	return nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the current size of the stack
func (s *Stack) Size() int {
	return len(*s)
}
