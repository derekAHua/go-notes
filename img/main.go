package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{Y: 255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Cos(s)*size/2
		pic.SetGray(x, int(y), color.Gray{})
	}

	fd, _ := os.Create("img/cos.png")
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)
	_ = png.Encode(fd, pic)
}
