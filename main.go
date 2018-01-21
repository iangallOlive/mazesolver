package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"time"
)

// Bash: clear && go build -o binary.exe && binary.exe maze1
// PowerShell (integrated VSCode terminal): clear; go build -o binary.exe; ./binary.exe maze1

func main() {
	if len(os.Args) < 2 {
		log.Println("No path to maze image")
		log.Println("Exiting...")
		os.Exit(1)
	}

	// Prepare solution folders if not exists

	if _, err := os.Stat("./solutions"); os.IsNotExist(err) {
		fError := os.Mkdir("solutions", os.ModePerm)
		if fError != nil {
			log.Fatal("Setup error:", fError)
		}
	}

	// Step 1: Load the image file
	path := os.Args[1]
	imagePath := fmt.Sprintf("./mazes/%v.png", path)
	fReader, err := os.Open(imagePath)
	defer fReader.Close()
	if err != nil {
		log.Fatalf("Error opening file. Error: %v", err)
	}

	// Step 2: Convert image into a mesh of Nodes
	img, _, err := image.Decode(fReader)
	if err != nil {
		log.Fatalf("Error decoding file. Error: %v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	log.Printf("Image is %v x %v", width, height)

	var nodes []Node
	NodeID := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// log.Println("x, y:", x, y)
			NodeID++
			var node Node
			tempR, tempG, tempB, _ := img.At(x, y).RGBA()             //uint32 values (Skip alpha value)
			r, g, b := int(tempR/257), int(tempG/257), int(tempB/257) // Somehow convert uint32 to 0-255 RGB color scale (Stackoverflow is great)
			if r == 255 && g == 255 && b == 255 {
				node = Node{isWall: false}
			} else {
				node = Node{isWall: true}
			}
			nodes = append(nodes, node)
		}
	}

	for i := range nodes {
		if i > 0 {
			nodes[i].left = &nodes[i-1]
		}

		if i+1 < len(nodes)-1 {
			nodes[i].right = &nodes[i+1]
		}

		if i-width > 0 {
			nodes[i].up = &nodes[i-width]
		}

		if i+width < len(nodes)-1 {
			nodes[i].down = &nodes[i+width]
		}

		nodes[i].checked = false
	}

	var potentialPoints []Point
	// Check top
	for i := 0; i < width; i++ {
		if nodes[i].isWall == false {
			potentialPoints = append(potentialPoints, Point{topSide: true, node: &nodes[i]})
		}
	}

	// Check bottom
	for i := (len(nodes) - 1) - width; i < len(nodes)-1; i++ {
		if nodes[i].isWall == false {
			potentialPoints = append(potentialPoints, Point{bottomSide: true, node: &nodes[i]})
		}
	}

	// Check left
	for i := (width - 1) + 1; i < len(nodes); i += width {
		if nodes[i].isWall == false {
			potentialPoints = append(potentialPoints, Point{leftSide: true, node: &nodes[i]})
		}
	}

	// Check right
	for i := width - 1; i <= len(nodes)-1; i += width {
		if nodes[i].isWall == false {
			potentialPoints = append(potentialPoints, Point{topSide: true, node: &nodes[i]})
		}
	}

	if len(potentialPoints) > 2 {
		log.Fatalf("Error: Too many potential starting points. Found %v starting points", len(potentialPoints))
	}
	log.Println("Total potential Points: ", len(potentialPoints))

	startPos := potentialPoints[0].node
	endPos := potentialPoints[1].node

	correspondingStartNode := make(map[string]string)
	correspondingStartNode["up"] = "down"
	correspondingStartNode["down"] = "up"
	correspondingStartNode["left"] = "right"
	correspondingStartNode["right"] = "left"

	stack := NewQueue()

	startPos.checked = true
	startPos.isSolution = true

	if potentialPoints[0].bottomSide == true {
		// Go up
		startPos.up.parent = startPos
		stack.Push(startPos.up)
	} else if potentialPoints[0].topSide == true {
		// Go down
		startPos.down.parent = startPos
		stack.Push(startPos.down)
	} else if potentialPoints[0].leftSide == true {
		// Go right
		startPos.right.parent = startPos
		stack.Push(startPos.right)
	} else {
		// Go left
		startPos.left.parent = startPos
		stack.Push(startPos.left)
	}

	start := time.Now()
	for {
		n := stack.Pop()
		if n == nil {
			// End of the stack, either errored or finished
			log.Println("Out of items")
			return
		}
		n.checked = true // Make sure current item wont be processed again

		if n == endPos {
			log.Printf("Solved after %v ms \n", time.Since(start).Nanoseconds()/int64(time.Millisecond))
			n.isSolution = true
			for node := endPos.parent; node.parent != nil; node = node.parent {
				node.isSolution = true
			}
			str, err := OutputSolution(path, height, width, nodes)
			if err != nil {
				log.Fatal("Error1:", err)
			}
			log.Println("OutputSolution returned:", str)
			return
		}

		if n.down != nil {
			if n.down.checked == false {
				if n.down.isWall == false {
					n.down.parent = n
					stack.Push(n.down)
				}
			}
		}

		if n.right != nil {
			if n.right.checked == false {
				if n.right.isWall == false {
					n.right.parent = n
					stack.Push(n.right)
				}
			}
		}

		if n.left != nil {
			if n.left.checked == false {
				if n.left.isWall == false {
					n.left.parent = n
					stack.Push(n.left)
				}
			}
		}

		if n.up != nil {
			if n.up.checked == false {
				if n.up.isWall == false {
					n.up.parent = n
					stack.Push(n.up)
				}
			}
		}
	}
}
