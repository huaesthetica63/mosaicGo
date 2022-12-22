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

func (i *Image) ResizeImage(new_width, new_height int) Image {
	var res Image
	res.width = new_width
	res.height = new_height
	res.pixels = make([][]ColorPixel, 0, new_height)
	for y := 0; y < new_height; y++ {
		row := make([]ColorPixel, 0, new_width)
		for x := 0; x < new_width; x++ {
			y_new := int(float32(i.height*y) / float32(new_height))
			x_new := int(float32(i.width*x) / float32(new_width))
			row = append(row, i.pixels[y_new][x_new])
		}
		res.pixels = append(res.pixels, row)
	}
	return res
}
func GetColorPixel(col color.Color) ColorPixel {
	r, g, b, a := col.RGBA()
	return ColorPixel{
		R: int(r / 255),
		G: int(g / 255),
		B: int(b / 255),
		A: int(a / 255),
	}
}
func (i *Image) Binarize() image.Image {
	middle := 255.0 / 2.0
	res := image.NewGray(image.Rect(0, 0, i.width, i.height))
	gray := i.ToGrayscale()
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			if gray.GrayAt(x, y).Y >= uint8(middle) {
				res.SetGray(x, y, color.Gray{Y: 255})
			} else {
				res.SetGray(x, y, color.Gray{Y: 0})
			}
		}
	}
	return res
}
func (i *Image) Mosaic(n uint8) image.Image {
	diapasone := 255.0 / float32(n)
	res := image.NewGray(image.Rect(0, 0, i.width, i.height))
	gray := i.ToGrayscale()
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			num_col := uint8(gray.GrayAt(x, y).Y / uint8(diapasone))
			var col uint8
			if num_col == 0 {
				col = 0
			} else if num_col == n-1 || num_col == 255 {
				col = 255
			} else {
				col = num_col*uint8(diapasone) + uint8(diapasone)/2
			}
			res.SetGray(x, y, color.Gray{Y: col})
		}
	}
	return res
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
func (i *Image) ToGrayscale() image.Gray {
	res := image.NewGray(image.Rect(0, 0, i.width, i.height))
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			res.Set(x, y, color.Gray{Y: uint8(RGBtoGray(i.pixels[y][x]))})
		}
	}
	return *res
}
func (i *Image) SaveGrayscaleToPng(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	gray := i.ToGrayscale()
	if err := png.Encode(f, &gray); err != nil {
		return err
	}
	return nil
}
func (i *Image) SaveBinarizeToPng(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, i.Binarize()); err != nil {
		return err
	}
	return nil
}
func (i *Image) SaveMosaicToPng(n uint8, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, i.Mosaic(n)); err != nil {
		return err
	}
	return nil
}
