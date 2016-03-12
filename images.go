// https://tour.golang.org/methods/25

package tour

import (
	"image"
	"image/color"
)

type Image struct{ X, Y int }

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.X, i.Y)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}
