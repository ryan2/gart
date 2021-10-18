package gart

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
)

func GeneratePNG(w io.Writer, height, width int) {
	skipw := 3
	skiph := 3
	colors := 255
	fmt.Println(skipw, skiph, colors)
	rect := image.Rect(0, 0, height, width)
	palette := make([]color.Color, 0, height*width)
	for i := 0; i < colors; i++ {
		palette = append(palette, getRandColor())
	}
	img := image.NewPaletted(rect, palette)
	for i := 0; i <= height; i += skipw {
		for j := 0; j <= width; j += skiph {
			idx := uint8(i+j) % uint8(colors)
			if i+j > (height+width)/3 {
				idx = uint8(i) % uint8(colors)
			}
			img.SetColorIndex(i, j, idx)
			skiph += 1
			skiph %= 3
		}
		skipw += 1
		skipw %= 3
	}
	png.Encode(w, img)
}
