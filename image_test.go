package main

import (
	"image"
	"image/color"
	"testing"
)

// createMockImage creates a simple test image with a specified pattern
func createMockImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Create a simple checkerboard pattern
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Alternate black and white pixels
			if (x+y)%2 == 0 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	return img
}

func TestProcessImage(t *testing.T) {
	// Create a small test image (4x4 pixels)
	img := createMockImage(4, 4)

	// Process the image
	data, width, height := processImage(img)

	// Check dimensions
	if width != 4 {
		t.Errorf("Expected width to be 4, got %d", width)
	}
	if height != 4 {
		t.Errorf("Expected height to be 4, got %d", height)
	}

	// For a 4x4 checkerboard image, we expect each row to be 0xAA (10101010 in binary)
	// Since we're packing 8 bits per byte, and we have 4 pixels per row,
	// we should have 4 bytes (one for each row, with 4 bits used in each byte)
	if len(data) != 4 {
		t.Errorf("Expected data length to be 4, got %d", len(data))
	}

	// Check the pattern (each row should be 0xA0 - 10100000 in binary)
	// The last 4 bits are unused and set to 0
	expectedPattern := byte(0xA0)
	for i, b := range data {
		if b != expectedPattern {
			t.Errorf("Row %d: Expected 0x%02X, got 0x%02X", i, expectedPattern, b)
		}
	}
}

// TestProcessImageWithOddDimensions tests the processImage function with an image that has odd dimensions
func TestProcessImageWithOddDimensions(t *testing.T) {
	// Create a test image with odd dimensions (5x3 pixels)
	img := createMockImage(5, 3)

	// Process the image
	data, width, height := processImage(img)

	// Check dimensions
	if width != 5 {
		t.Errorf("Expected width to be 5, got %d", width)
	}
	if height != 3 {
		t.Errorf("Expected height to be 3, got %d", height)
	}

	// For a 5x3 image, we expect 3 bytes (one for each row)
	// Each row will have 5 bits used, with 3 bits of padding
	if len(data) != 3 {
		t.Errorf("Expected data length to be 3, got %d", len(data))
	}

	// Check the pattern for each row
	// For a 5-pixel row with alternating 1s and 0s starting with 1, we expect 0xA8 (10101000 in binary)
	expectedPattern := byte(0xA8)
	for i, b := range data {
		if b != expectedPattern {
			t.Errorf("Row %d: Expected 0x%02X, got 0x%02X", i, expectedPattern, b)
		}
	}
}

// TestProcessImageEmpty tests the processImage function with an empty image (0x0)
func TestProcessImageEmpty(t *testing.T) {
	// Create an empty image (0x0 pixels)
	img := createMockImage(0, 0)

	// Process the image
	data, width, height := processImage(img)

	// Check dimensions
	if width != 0 {
		t.Errorf("Expected width to be 0, got %d", width)
	}
	if height != 0 {
		t.Errorf("Expected height to be 0, got %d", height)
	}

	// For an empty image, we expect an empty byte array
	if len(data) != 0 {
		t.Errorf("Expected data length to be 0, got %d", len(data))
	}
}
