package color_mosaic

import (
	"image"
	"image/color"
	"main/image_processing"
)

//interface for creating different types of mosaics
type Mosaic interface {
	MakeMosaic(i image_processing.Image) image.Image
}
type GrayscaleMosaic struct {
	n uint8 //n different grayscale colors
}

func NewGrayscaleMosaic(num_col uint8) GrayscaleMosaic {
	return GrayscaleMosaic{n: num_col}
}
func (gr GrayscaleMosaic) MakeMosaic(i image_processing.Image) image.Image {
	diapasone := 255.0 / float32(gr.n)
	res := image.NewGray(image.Rect(0, 0, i.Width, i.Height))
	gray := i.ToGrayscale()
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			num_col := uint8(gray.GrayAt(x, y).Y / uint8(diapasone))
			var col uint8
			//first color = 0
			if num_col == 0 {
				col = 0
				//last color = 255
			} else if num_col == gr.n-1 || num_col == 255 {
				col = 255
			} else {
				col = num_col*uint8(diapasone) + uint8(diapasone)/2
			}
			res.SetGray(x, y, color.Gray{Y: col})
		}
	}
	return res
}

type ColorMosaic struct {
	palette []color.Color
}

func NewBlueMosaic() ColorMosaic {
	colors := make([]color.Color, 3)
	colors[2] = color.RGBA{R: 204, G: 229, B: 255, A: 255}
	colors[1] = color.RGBA{R: 0, G: 128, B: 255, A: 255}
	colors[0] = color.RGBA{R: 0, G: 51, B: 102, A: 255}
	return ColorMosaic{palette: colors}
}
func NewGreenMosaic() ColorMosaic {
	colors := make([]color.Color, 3)
	colors[2] = color.RGBA{R: 204, G: 255, B: 229, A: 255}
	colors[1] = color.RGBA{R: 0, G: 204, B: 102, A: 255}
	colors[0] = color.RGBA{R: 0, G: 102, B: 51, A: 255}
	return ColorMosaic{palette: colors}
}
func NewRedMosaic() ColorMosaic {
	colors := make([]color.Color, 3)
	colors[2] = color.RGBA{R: 255, G: 204, B: 255, A: 255}
	colors[1] = color.RGBA{R: 255, G: 51, B: 51, A: 255}
	colors[0] = color.RGBA{R: 102, G: 0, B: 0, A: 255}
	return ColorMosaic{palette: colors}
}
func NewPeachMosaic() ColorMosaic {
	colors := make([]color.Color, 3)
	colors[2] = color.RGBA{R: 255, G: 229, B: 204, A: 255}
	colors[1] = color.RGBA{R: 255, G: 153, B: 51, A: 255}
	colors[0] = color.RGBA{R: 153, G: 76, B: 0, A: 255}
	return ColorMosaic{palette: colors}
}
func (cl ColorMosaic) MakeMosaic(i image_processing.Image) image.Image {
	n := uint8(len(cl.palette))
	diapasone := 255.0 / float32(n)
	res := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))
	gray := i.ToGrayscale()
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			num_col := uint8(gray.GrayAt(x, y).Y / uint8(diapasone))
			var col color.Color
			if num_col == n {
				num_col = n - 1
			}
			col = cl.palette[num_col]
			res.Set(x, y, col)
		}
	}
	return res
}
