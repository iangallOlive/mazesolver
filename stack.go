package main

// Stack : LIFO stack
type Stack struct {
	count int
	nodes []*Node
}

// NewStack : Returns a new Stack type
func NewStack() *Stack {
	return &Stack{}
}

// Push : Adds the element to the Stack
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop : Slices last element of the Stack and returns it
func (s *Stack) Pop() *Node {
	if s.count > 0 {
		s.count--
		return s.nodes[s.count]
	}
	return nil
}

// Count : Returns the current count of the Stack
func (s *Stack) Count() int {
	return s.count
}
