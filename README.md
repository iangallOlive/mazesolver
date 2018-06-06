# Maze Solver

A tool in Golang, built to solve mazes using [Depth First Search](https://en.wikipedia.org/wiki/Depth-first_search).

The algorithm can be switched up for a [Breadth-first search](https://en.wikipedia.org/wiki/Breadth-first_search). Check out [Algorithms](#algorithms)

## Info

This project was developed to learn graph theory & algorithms.

This project is not intended to be used for anything remotely serious. Bug fixes & features will be added as I please.


## How it works
The process loads in an image in the _/mazes_ folder. It then starts creating nodes with properties, based on the color of the pixel.

Black pixels are walls, white pixels are paths.

White pixels on the edge of the picture are points. There can only be 2 points, and the program then chooses one of them as starting, and the other the end point.

It then connects all the nodes as they were connected in the picture, by giving them a up, down, left, right pointer, pointing to that neighbour node.

Take the point designated as start, and the point designated as end, create a stack & and run a Depth First Search, looking for that end point.

## Installation

Since this only relies on stdlib, no dependencies will be installed.

``` go get github.com/jkob/mazesolver ```


## Algorithms

In order to switch to a BFS, you will have to create a queue instead of a stack. Please note that this will severely increase the time it takes to solve huge mazes (50k x 50k).

I recommend leaving one commented out, if you want to ever switch back without pain.

Locate where the variable **stack** is initialized and assigned. (_main.go_)

Instead of ```stack := NewStack()```, assign it to ```NewQueue()```

