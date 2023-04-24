package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawDictionary(dict *Dictionary2) {
	for _, achar := range dict.Character {
		drawChar2(&achar)
	}
}

func drawChar2(onechar *Character2) {
	filename := fmt.Sprintf("(%s).png", onechar.Utf8)
	//fmt.Printf("Filename: %s\n", filename)
	const width = 1000
	const height = 1000

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x0, 0x0, 0x0, 0xff})
	gc.SetStrokeColor(color.RGBA{0xe0, 0xe0, 0xe0, 0xff})
	gc.SetLineWidth(10)

	for _, onestroke := range onechar.Strokes.Stroke {
		drawOneStroke(gc, &onestroke)
	}

	// Save to file
	draw2dimg.SaveToPngFile(filename, dest)
}

func drawOneStroke(gc *draw2dimg.GraphicContext, sk *Stroke2) {
	gc.BeginPath() // Initialize a new path
	if len(sk.Point) <= 1 {
		return
	}

	firstPt := sk.Point[0]
	gc.MoveTo(str2f64(firstPt.X), str2f64(firstPt.Y))
	for i := 1; i < len(sk.Point); i++ {
		pt := sk.Point[i]
		gc.LineTo(str2f64(pt.X), str2f64(pt.Y))
	}
	gc.Stroke()
	gc.Close()
}
