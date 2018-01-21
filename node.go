package main

// Node : Represents a pixel on the image
type Node struct {
	up         *Node
	down       *Node
	left       *Node
	right      *Node
	isWall     bool
	checked    bool
	parent     *Node
	isSolution bool
}
