package render

import (
	"bytes"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"labl/pkg/templates"
	"log/slog"
	"math"

	"github.com/go-pdf/fpdf"
)

type Image struct {
	Name             string
	Center           bool
	CenterVertical   bool
	CenterHorizontal bool
	Position         templates.Position
	Size             templates.Size
	Color            [3]int
	Resources        *Resources
}

func (i Image) Render(pdf *fpdf.Fpdf, pos RenderPosition) {
	x, y := Percent(pos.Width, float64(i.Position.X))+pos.X, Percent(pos.Height, float64(i.Position.Y))+pos.Y
	containerWidth, containerHeight := Percent(pos.Width, float64(i.Size.Width)), Percent(pos.Height, float64(i.Size.Height))

	if i.Center {
		x = x + (pos.Width-containerWidth)/2
		y = y + (pos.Height-containerHeight)/2
	}
	if i.CenterVertical {
		y = y + (pos.Height-containerHeight)/2
	}
	if i.CenterHorizontal {
		x = x + (pos.Width-containerWidth)/2
	}

	info := pdf.GetImageInfo(i.Name)
	imageWidth, imageHeight := float64(info.Width()), float64(info.Height())
	aspectRatio := imageWidth / imageHeight

	// Fit the image to the container while maintaining the aspect ratio
	var width, height float64
	if containerWidth/containerHeight > aspectRatio {
		height = containerHeight
		width = height * aspectRatio
	} else {
		width = containerWidth
		height = width / aspectRatio
	}

	// Center the image in the container
	x = x + (containerWidth-width)/2
	y = y + (containerHeight-height)/2

	x = math.Max(0, x)
	y = math.Max(0, y)

	imageName := i.Name
	if i.Color != [3]int{} {
		imageName += randSeq(10)
		data, err := i.Resources.GetBytes(i.Name)
		if err != nil {
			slog.Error("error while colorizing image", "err", err, "len", len(data))
			return
		}

		img, err := png.Decode(bytes.NewBuffer(data))
		if err != nil {
			slog.Error("error while colorizing image", "err", err, "len", len(data))
			return
		}

		c := color.RGBA{uint8(i.Color[0]), uint8(i.Color[1]), uint8(i.Color[2]), 255}
		img = colorizeImage(c, img)
		var encoded bytes.Buffer
		err = png.Encode(&encoded, img)
		if err != nil {
			slog.Error("error while colorizing image", "err", err)
			return
		}

		pdf.RegisterImageOptionsReader(imageName, fpdf.ImageOptions{
			ImageType: "PNG",
			ReadDpi:   true,
		}, &encoded)
	}

	pdf.ImageOptions(imageName, x, y, width, height, false, fpdf.ImageOptions{}, 0, "")
}
