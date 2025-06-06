package main

import (
	"testing"
)

func TestTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "Multiple words",
			input:    "hello world",
			expected: "Hello World",
		},
		{
			name:     "With special characters",
			input:    "hello-world",
			expected: "Hello World",
		},
		{
			name:     "With numbers",
			input:    "hello123 world456",
			expected: "Hello123 World456",
		},
		{
			name:     "Already title case",
			input:    "Hello World",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := titleCase(tt.input)
			if result != tt.expected {
				t.Errorf("titleCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPNGFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "Valid PNG file",
			path:     "image.png",
			expected: true,
		},
		{
			name:     "Valid PNG file with uppercase extension",
			path:     "image.PNG",
			expected: true,
		},
		{
			name:     "Valid PNG file with mixed case extension",
			path:     "image.PnG",
			expected: true,
		},
		{
			name:     "Invalid file extension",
			path:     "image.jpg",
			expected: false,
		},
		{
			name:     "No file extension",
			path:     "image",
			expected: false,
		},
		{
			name:     "Path with directory",
			path:     "/path/to/image.png",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPNGFile(tt.path)
			if result != tt.expected {
				t.Errorf("isPNGFile(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestIsGoFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "Valid Go file",
			path:     "file.go",
			expected: true,
		},
		{
			name:     "Valid Go file with uppercase extension",
			path:     "file.GO",
			expected: true,
		},
		{
			name:     "Valid Go file with mixed case extension",
			path:     "file.Go",
			expected: true,
		},
		{
			name:     "Invalid file extension",
			path:     "file.txt",
			expected: false,
		},
		{
			name:     "No file extension",
			path:     "file",
			expected: false,
		},
		{
			name:     "Path with directory",
			path:     "/path/to/file.go",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isGoFile(tt.path)
			if result != tt.expected {
				t.Errorf("isGoFile(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}
