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

// main is the entry point of the program. It processes command-line arguments,
// reads the input PNG file, converts it to a byte array, and writes the result to a Go file.
func main() {
	// Check if the required command-line arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . input.png output.go")
		return
	}

	// Extract input and output file paths from command-line arguments
	inputPath := os.Args[1]
	outputPath := os.Args[2]

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
	data, width, height := processImage(img)
	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	// Generate the output Go file
	err = generateGoFile(outputPath, varName, data, width, height)
	if err != nil {
		panic(err)
	}

	// Print a success message
	fmt.Printf("Done. Bytes written to %s\n", outputPath)
}
