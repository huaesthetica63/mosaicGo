package image_processing

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type ColorPixel struct {
	R int
	G int
	B int
	A int
}
type Image struct {
	pixels [][]ColorPixel
	width  int
	height int
}

func GetColorPixel(col color.Color) ColorPixel {
	r, g, b, a := col.RGBA()
	return ColorPixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257),
	}
}
func (i *Image) LoadImage(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	img, fileformat, err := image.Decode(file)
	if err != nil {
		return err
	}
	fmt.Println(fileformat)
	bounds := img.Bounds()
	i.height = bounds.Max.Y
	i.width = bounds.Max.X
	i.pixels = make([][]ColorPixel, 0, i.height)
	for y := 0; y < i.height; y++ {
		row := make([]ColorPixel, 0, i.width)
		for x := 0; x < i.width; x++ {
			row = append(row, GetColorPixel(img.At(x, y)))
		}
		i.pixels = append(i.pixels, row)
	}
	return nil
}
func RGBtoGray(pix ColorPixel) float32 {
	x := 0.3*float32(pix.R) + 0.59*float32(pix.G) + 0.11*float32(pix.B)
	return x
}
func (i *Image) ToGrayscale() image.Image {
	res := image.NewGray(image.Rect(0, 0, i.width, i.height))
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			res.Set(x, y, color.Gray{Y: uint8(RGBtoGray(i.pixels[y][x]))})
		}
	}
	return res
}
func (i *Image) SaveGrayscaleToPng(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, i.ToGrayscale()); err != nil {
		return err
	}
	return nil
}
