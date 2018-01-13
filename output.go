package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
)

// OutputSolution : Renders the maze image and solution, side by side
func OutputSolution(mazeName string, imageHeight, imageWidth int) (string, error) {

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

	fmt.Println("mazeImage(0,0):", mazeImage.At(0, 0))

	pixels := make(map[string]color.RGBA)
	gap := 3
	totalWidth := imageWidth*2 + gap
	// Pixel colors
	pixels["red"] = color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
	pixels["black"] = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
	pixels["white"] = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
	pixels["opaque"] = color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(0)}

	image := image.NewRGBA(image.Rect(0, 0, totalWidth, imageHeight))

	for h := 0; h < imageHeight; h++ {
		for w := 0; w < totalWidth; w++ {
			if w < imageWidth {
				// Copy phase
				mazePixel := mazeImage.At(w, h).RGBA()
				fmt.Println("mazePixel:", mazePixel)
				image.SetRGBA(w, h, mazePixel)
			}

			if w > imageWidth && w < totalWidth {
				// Gap phase
				image.SetRGBA(w, h, pixels["black"])
			}

			if w > imageWidth+gap {
				// Solution phase
				image.SetRGBA(w, h, pixels["red"])
			}
		}
	}

	return "", nil
}
