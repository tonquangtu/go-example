package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	bound image.Rectangle
	color color.Model
	data  [][]color.Color
}

func newImage(width, height int) *Image {
	img := &Image{bound: image.Rect(0, 0, width, height), color: color.RGBAModel}
	img.Init()
	return img
}

func (img *Image) Init() {
	width, height := img.bound.Max.X, img.bound.Max.Y
	img.data = make([][]color.Color, height)
	for i := 0; i < height; i++ {
		img.data[i] = make([]color.Color, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.data[i][j] = color.RGBA{R: uint8(i * j), G: uint8(i * j), B: 255, A: 255}
		}
	}
}

func (img *Image) Bounds() image.Rectangle {
	return img.bound
}

func (img *Image) At(x, y int) color.Color {
	return img.data[y][x]
}

func (img *Image) ColorModel() color.Model {
	return img.color
}

func main() {
	pic.ShowImage(newImage(1000, 1000))
}
