package main

// Point : Starting/end point for the maze.
type Point struct {
	topSide    bool
	leftSide   bool
	rightSide  bool
	bottomSide bool
	node       *Node
}
