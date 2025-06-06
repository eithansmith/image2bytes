package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateGoFile(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "image2bytes_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func(path string) {
		err = os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to remove temp dir: %v", err)
		}
	}(tempDir) // Clean up after the test

	// Create a test output file path
	outputPath := filepath.Join(tempDir, "test_output.go")

	// Create test data
	data := []byte{0xA0, 0xA0, 0xA0, 0xA0}
	width, height := 4, 4
	varName := "TestImage"

	// Generate the Go file
	err = generateGoFile(outputPath, varName, data, width, height)
	if err != nil {
		t.Fatalf("generateGoFile failed: %v", err)
	}

	// Read the generated file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Check if the file contains the expected content
	expectedContent := []string{
		"package main",
		"// TestImageWidth and TestImageHeight define image dimensions",
		"const TestImageWidth = 4",
		"const TestImageHeight = 4",
		"var TestImage = []byte{",
		"0xA0, 0xA0, 0xA0, 0xA0,",
	}

	for _, expected := range expectedContent {
		if !bytes.Contains(content, []byte(expected)) {
			t.Errorf("Generated file does not contain expected content: %s", expected)
		}
	}
}

// TestGenerateGoFileError tests the error handling in generateGoFile
func TestGenerateGoFileError(t *testing.T) {
	// Create an invalid output path (in a non-existent directory)
	outputPath := "/nonexistent/directory/test_output.go"

	// Create test data
	data := []byte{0xA0, 0xA0, 0xA0, 0xA0}
	width, height := 4, 4
	varName := "TestImage"

	// Generate the Go file, which should fail
	err := generateGoFile(outputPath, varName, data, width, height)

	// Check that an error was returned
	if err == nil {
		t.Errorf("Expected an error when creating a file in a non-existent directory, but got nil")
	}
}

// TestGenerateGoFileWithEmptyData tests generateGoFile with empty data
func TestGenerateGoFileWithEmptyData(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "image2bytes_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to remove temp dir: %v", err)
		}
	}(tempDir)

	// Create a test output file path
	outputPath := filepath.Join(tempDir, "empty_output.go")

	// Create empty test data
	var data []byte
	width, height := 0, 0
	varName := "EmptyImage"

	// Generate the Go file
	err = generateGoFile(outputPath, varName, data, width, height)
	if err != nil {
		t.Fatalf("generateGoFile failed: %v", err)
	}

	// Read the generated file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Check if the file contains the expected content
	expectedContent := []string{
		"package main",
		"// EmptyImageWidth and EmptyImageHeight define image dimensions",
		"const EmptyImageWidth = 0",
		"const EmptyImageHeight = 0",
		"var EmptyImage = []byte{",
		"}",
	}

	for _, expected := range expectedContent {
		if !bytes.Contains(content, []byte(expected)) {
			t.Errorf("Generated file does not contain expected content: %s", expected)
		}
	}
}
