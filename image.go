package main

import (
	"image"

	"golang.org/x/image/draw"
)

// processImage converts any image to packed 1bpp bytes (MSB-first).
// Returns (data, width, height).
func processImage(src image.Image) ([]byte, int, int, error) {
	const targetW, targetH = 296, 128 // Badger 2040W

	// 1) Resize to panel resolution (preserve aspect to fill; adjust if you prefer letterboxing)
	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	// 2) Convert to 1bpp and pack bits: 1 = black, 0 = white (invert if your driver expects opposite)
	// Luminance threshold ~50% (tweak if needed)
	const thresh = 0x8000

	data := make([]byte, 0, (targetW*targetH+7)/8)
	for y := 0; y < targetH; y++ {
		var byteAcc uint8
		bitCount := 0
		for x := 0; x < targetW; x++ {
			r, g, b, _ := dst.At(x, y).RGBA() // 16-bit per channel (0..65535)

			// Perceptual luminance (ITU-R BT.601-ish), scaled to 0..65535
			luma := (299*r + 587*g + 114*b) / 1000

			var bit uint8
			if luma < thresh {
				bit = 1 // darker pixel -> black
			} else {
				bit = 0 // lighter pixel -> white
			}

			byteAcc = (byteAcc << 1) | bit
			bitCount++

			if bitCount == 8 {
				data = append(data, byteAcc)
				byteAcc, bitCount = 0, 0
			}
		}
		// Pad any partial byte at end of row
		if bitCount > 0 {
			byteAcc <<= (8 - bitCount)
			data = append(data, byteAcc)
		}
	}

	return data, targetW, targetH, nil
}
