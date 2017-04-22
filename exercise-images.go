package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) / 2), uint8((x + y) / 2), 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
