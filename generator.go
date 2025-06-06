package main

import (
	"fmt"
	"os"
)

// generateGoFile writes the byte array data to a Go file
func generateGoFile(outputPath string, varName string, data []byte, width, height int) error {
	// Create the output Go file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
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

	return nil
}
