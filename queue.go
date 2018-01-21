package main

// Queue : Is a simple Queue implementation
type Queue struct {
	count int
	n     []*Node
}

// NewQueue : Returns a Queue
func NewQueue() *Queue {
	return &Queue{}
}

// Push : Appends item to Queue
func (q *Queue) Push(n *Node) {
	q.n = append(q.n, n)
	q.count++
}

// Pop : Pops item from Queue
func (q *Queue) Pop() *Node {
	if q.count > 0 {
		q.count--
		node := q.n[0]
		q.n = q.n[1:]
		return node
	}

	return nil
}

// Count : Return items currently in Queue
func (q *Queue) Count() int {
	return q.count
}
