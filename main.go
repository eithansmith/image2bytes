package main

// image2bytes is a utility that converts PNG images to byte arrays for embedding in Go code.
// It processes the image pixel by pixel, converting it to a monochrome representation
// where each bit represents a pixel (1 for black, 0 for white).

import (
	"fmt"
	"image/png"
	"os"
	"strings"
)

const (
	inputPath  = "input.png"
	outputPath = "output.go"
)

// main is the entry point of the program. It processes command-line arguments,
// reads the input PNG file, converts it to a byte array, and writes the result to a Go file.
func main() {
	// Validate that inputPath is a PNG file
	if !isPNGFile(inputPath) {
		fmt.Println("Error: Input file must be a PNG file (with .png extension)")
		return
	}

	// Validate that outputPath is a Go file
	if !isGoFile(outputPath) {
		fmt.Println("Error: Output file must be a Go file (with .go extension)")
		return
	}

	// Generate a variable name for the output Go file based on the output file name
	varName := strings.TrimSuffix(titleCase(strings.ReplaceAll(outputPath, ".go", "")), ".go")

	// Open the input PNG file
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	// Ensure the file is closed when the function returns
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Decode the PNG image
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// Process the image
	data, width, height, err := processImage(img)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	// Generate the output Go file
	err = generateGoFile(outputPath, varName, data, width, height)
	if err != nil {
		panic(err)
	}

	// Print a success message
	fmt.Printf("Done. Bytes written to %s\n", outputPath)
}
