package gart

import (
	"image/color"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandColor() color.Color {
	r := uint8(rand.Int31n(255))
	g := uint8(rand.Int31n(255))
	b := uint8(rand.Int31n(255))
	a := uint8(rand.Int31n(255))
	return color.RGBA{r, g, b, a}
}
