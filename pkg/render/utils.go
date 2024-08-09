package render

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
)

func Percent(value float64, percent float64) float64 {
	return value * percent / 100
}

type RenderPosition struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func HexToColor(hex string) (color.Color, error) {
	var c color.RGBA
	if _, err := fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B); err != nil {
		return nil, err
	}
	c.A = 255
	return c, nil
}

func colorizeImage(c color.Color, img image.Image) image.Image {
	bounds := img.Bounds()
	coloredImg := image.NewRGBA(bounds)

	// Extract the red, green, and blue components of the color
	r, g, b, _ := c.RGBA()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			origR, origG, origB, origA := originalColor.RGBA()

			// Apply the color filter
			newColor := color.RGBA{
				R: uint8((origR * r) / math.MaxUint16 / 256),
				G: uint8((origG * g) / math.MaxUint16 / 256),
				B: uint8((origB * b) / math.MaxUint16 / 256),
				A: uint8(origA / 256),
			}

			coloredImg.Set(x, y, newColor)
		}
	}

	return coloredImg
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
