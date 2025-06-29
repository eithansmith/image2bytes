# image2bytes

A utility that converts PNG images to byte arrays for embedding in Go code. It processes the image pixel by pixel, converting it to a monochrome representation where each bit represents a pixel (1 for black, 0 for white).

## Features

- Converts PNG images to Go byte arrays
- Automatically generates constants for image dimensions
- Creates ready-to-use Go code files
- Optimized for monochrome images
- Simple command-line interface

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/image2bytes.git

# Navigate to the project directory
cd image2bytes

# Build the project
go build
```

Alternatively, you can install directly using `go get`:

```bash
go get github.com/yourusername/image2bytes
```

## Usage

```bash
# Basic usage
go run . input.png output.go

# Or if you've built the binary
./image2bytes input.png output.go
```

### Example

Convert the included gopher.png to a Go byte array:

```bash
go run . input.png output.go
```

This will generate a file named `output.go` containing:

```go
package main

// OutputWidth and OutputHeight define image dimensions
const OutputWidth = 296
const OutputHeight = 128

var Output = []byte{
    // Byte array data representing the image
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    // ... more bytes ...
}
```

## How It Works

1. The program reads a PNG image file
2. It converts the image to grayscale
3. Each pixel is converted to a bit (1 for black, 0 for white)
4. Bits are packed into bytes (8 bits per byte)
5. The bytes are formatted as a Go byte array in the output file
6. Constants for image dimensions are included in the output file

## Use Cases

- Embedding images in Go applications without external files
- Creating graphics for embedded systems or microcontrollers
- Generating bitmap fonts for displays
- Any application where you need to include image data directly in your code

## Requirements

- Go 1.24 or later

## Testing

The project includes a comprehensive test suite to ensure functionality and reliability. The tests cover:

- String processing functions (titleCase)
- File validation functions (isPNGFile, isGoFile)
- Image processing logic (processImage)
- File generation logic (generateGoFile)

### Running Tests

To run the tests:

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage report
go test -cover
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
