package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

// OutputSolution : Renders the maze image and solution, side by side
func OutputSolution(mazeName string, imageHeight, imageWidth int, nodes []Node) (string, error) {

	solutionName := fmt.Sprintf("solution_%v", mazeName)

	path := fmt.Sprintf("./solutions/%v.png", solutionName)
	solutionReader, err := os.Create(path)
	defer solutionReader.Close()
	if err != nil {
		log.Fatal("Creating solution file error:", err)
	}

	mazeReader, err := os.Open(fmt.Sprintf("mazes/%v.png", mazeName))
	if err != nil {
		log.Fatal("Opening maze file error:", err)
	}
	defer mazeReader.Close()

	mazeImage, _, err := image.Decode(mazeReader)
	if err != nil {
		return "", err
	}

	pixels := make(map[string]color.RGBA)
	gap := 3
	totalWidth := imageWidth*2 + gap
	// Pixel colors
	pixels["red"] = color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
	pixels["black"] = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
	pixels["white"] = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
	pixels["opaque"] = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(0)}
	pixels["purple"] = color.RGBA{R: uint8(102), G: uint8(51), B: uint8(153), A: uint8(255)}

	image := image.NewRGBA(image.Rect(0, 0, totalWidth, imageHeight))

	for h := 0; h < imageHeight; h++ {
		for w := 0; w < totalWidth; w++ {
			if w < imageWidth {
				// Copy phase
				r, g, b, a := mazeImage.At(w, h).RGBA()
				mazePixel := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
				image.SetRGBA(w, h, mazePixel)
			}

			if w >= imageWidth && w < imageWidth+gap {
				// Gap phase
				image.SetRGBA(w, h, pixels["opaque"])
			}

			if w >= (imageWidth + gap) {
				// Solution phase
				var pixelToSet color.RGBA
				nodeX := w - (imageWidth + gap)
				nodeY := h * imageHeight
				selectedNode := nodes[nodeY+nodeX]

				if selectedNode.isSolution == true {
					// log.Println("Is part of solution")
					pixelToSet = pixels["red"]
				} else if selectedNode.isWall == true {
					// log.Println("is wall")
					pixelToSet = pixels["black"]
				} else {
					// log.Println("is path")
					pixelToSet = pixels["white"]
				}

				image.SetRGBA(w, h, pixelToSet)
			}
		}
	}

	err = png.Encode(solutionReader, image)
	if err != nil {
		log.Fatal("Error encoding:", err)
	}

	return solutionName, nil
}
