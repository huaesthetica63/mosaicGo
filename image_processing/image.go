package image_processing

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
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
	Pixels [][]ColorPixel
	Width  int
	Height int
}

func (i *Image) ResizeImage(new_width, new_height int) Image {
	var res Image
	res.Width = new_width
	res.Height = new_height
	res.Pixels = make([][]ColorPixel, 0, new_height)
	for y := 0; y < new_height; y++ {
		row := make([]ColorPixel, 0, new_width)
		for x := 0; x < new_width; x++ {
			y_new := int(float32(i.Height*y) / float32(new_height))
			x_new := int(float32(i.Width*x) / float32(new_width))
			row = append(row, i.Pixels[y_new][x_new])
		}
		res.Pixels = append(res.Pixels, row)
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
func (i *Image) Binarize(threshold float32) image.Image {
	t := 255.0 * threshold
	res := image.NewGray(image.Rect(0, 0, i.Width, i.Height))
	gray := i.ToGrayscale()
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			if gray.GrayAt(x, y).Y >= uint8(t) {
				res.SetGray(x, y, color.Gray{Y: 255})
			} else {
				res.SetGray(x, y, color.Gray{Y: 0})
			}
		}
	}
	return res
}
func (i *Image) LoadImageBytes(data []byte) error {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}
	bounds := img.Bounds()
	i.Height = bounds.Max.Y
	i.Width = bounds.Max.X
	i.Pixels = make([][]ColorPixel, 0, i.Height)
	for y := 0; y < i.Height; y++ {
		row := make([]ColorPixel, 0, i.Width)
		for x := 0; x < i.Width; x++ {
			row = append(row, GetColorPixel(img.At(x, y)))
		}
		i.Pixels = append(i.Pixels, row)
	}
	return nil
}
func (i *Image) LoadImage(filename string) error {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	bounds := img.Bounds()
	i.Height = bounds.Max.Y
	i.Width = bounds.Max.X
	i.Pixels = make([][]ColorPixel, 0, i.Height)
	for y := 0; y < i.Height; y++ {
		row := make([]ColorPixel, 0, i.Width)
		for x := 0; x < i.Width; x++ {
			row = append(row, GetColorPixel(img.At(x, y)))
		}
		i.Pixels = append(i.Pixels, row)
	}
	return nil
}
func RGBtoGray(pix ColorPixel) float32 {
	x := 0.3*float32(pix.R) + 0.59*float32(pix.G) + 0.11*float32(pix.B)
	return x
}
func (i *Image) ToGrayscale() image.Gray {
	res := image.NewGray(image.Rect(0, 0, i.Width, i.Height))
	for y := 0; y < i.Height; y++ {
		for x := 0; x < i.Width; x++ {
			res.Set(x, y, color.Gray{Y: uint8(RGBtoGray(i.Pixels[y][x]))})
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
func SaveToPng(i image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, i); err != nil {
		return err
	}
	return nil
}
