package tour

import (
	"image"
	"image/color"
	"testing"
)

func TestImage(t *testing.T) {
	x := 7
	y := 9
	var i image.Image = Image{x, y}

	if i.Bounds() != image.Rect(0, 0, x, y) {
		t.Errorf("Image{%d, %d}.Bounds() == %v, want %v",
			x, y, i.Bounds(), image.Rect(0, 0, x, y))
	}

	if i.ColorModel() != color.RGBAModel {
		t.Errorf("Image{%d, %d}.ColorModel() == %v, want %v",
			x, y, i.ColorModel(), color.RGBAModel)
	}
}
