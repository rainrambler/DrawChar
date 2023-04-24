package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func drawChar(onechar *Character2) {
	filename := fmt.Sprintf("%s(%s).png", onechar.Text, onechar.Utf8)
	const width = 1000
	const height = 1000

	im := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})

	drawline(100, 100, 300, 300, drawDot(im))
	//newImg := image.NewRGBA

	outFile, err := os.Create(filename)
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, im)
	if err != nil {
		panic(err)
	}
	err = b.Flush()
	if err != nil {
		panic(err)
	}
}

// https://github.com/akavel/polyclip-go/blob/master/polyutil/draw.go
// Putpixel describes a function expected to draw a point on a bitmap at (x, y) coordinates.
type Putpixel func(x, y int)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: handle int overflow etc.
func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func drawDot(img draw.Image) func(x, y int) {
	return func(x, y int) {
		img.Set(x, y, &color.White)
	}
}
