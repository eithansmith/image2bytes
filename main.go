package main

// image2bytes is a utility that converts PNG images to byte arrays for embedding in Go code.
// It processes the image pixel by pixel, converting it to a monochrome representation
// where each bit represents a pixel (1 for black, 0 for white).

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strings"
	"unicode"
)

// titleCase returns a copy of the string s with the first letter of each word capitalized.
// This is a replacement for the deprecated strings.Title function.
func titleCase(s string) string {
	// Split the string into words
	words := strings.Fields(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return ' '
	}, s))

	// Capitalize the first letter of each word
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}

	// Join the words back together
	return strings.Join(words, "")
}

// main is the entry point of the program. It processes command-line arguments,
// reads the input PNG file, converts it to a byte array, and writes the result to a Go file.
func main() {
	// Check if the required command-line arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run image2bytes.go input.png output.go")
		return
	}

	// Extract input and output file paths from command-line arguments
	inputPath := os.Args[1]
	outputPath := os.Args[2]
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

	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	// Initialize the byte array to store the processed image data
	var data []byte
	// Process the image row by row
	for y := 0; y < height; y++ {
		var row byte = 0
		bitCount := 0
		// Process each pixel in the current row
		for x := 0; x < width; x++ {
			// Convert the pixel to grayscale
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			// Determine if the pixel is black (1) or white (0)
			// Pixels with brightness less than 128 are considered black
			bit := byte(0)
			if gray.Y < 128 {
				bit = 1
			}
			// Add the bit to the current byte (shift left and OR with the new bit)
			row = (row << 1) | bit
			bitCount++

			// When we have 8 bits (a complete byte), add it to the data array
			if bitCount == 8 {
				data = append(data, row)
				row = 0
				bitCount = 0
			}
		}
		// If we have remaining bits at the end of a row (not a complete byte),
		// pad with zeros and add to the data array
		if bitCount > 0 {
			row <<= 8 - bitCount
			data = append(data, row)
		}
	}

	// Create the output Go file
	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	// Ensure the output file is closed when the function returns
	defer func(outFile *os.File) {
		err = outFile.Close()
		if err != nil {
			panic(err)
		}
	}(outFile)

	// Write the Go code to the output file
	// Start with the package declaration
	_, _ = fmt.Fprintf(outFile, "package main\n\n")
	// Add constants for the image dimensions
	_, _ = fmt.Fprintf(outFile, "// %sWidth and %sHeight define image dimensions\n", varName, varName)
	_, _ = fmt.Fprintf(outFile, "const %sWidth = %d\n", varName, width)
	_, _ = fmt.Fprintf(outFile, "const %sHeight = %d\n\n", varName, height)
	// Begin the byte array declaration
	_, _ = fmt.Fprintf(outFile, "var %s = []byte{\n", varName)
	// Write the byte array data in a formatted way (12 bytes per line)
	for i, b := range data {
		if i%12 == 0 {
			_, _ = fmt.Fprintf(outFile, "\n\t")
		}
		_, _ = fmt.Fprintf(outFile, "0x%02X, ", b)
	}
	// Close the byte array declaration
	_, _ = fmt.Fprintf(outFile, "\n}\n")

	// Print a success message
	fmt.Printf("Done. Bytes written to %s\n", outputPath)
}
