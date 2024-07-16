package render

import (
	_ "image/jpeg"
	_ "image/png"
	"labl/pkg/templates"
	"math"

	"github.com/go-pdf/fpdf"
)

type Image struct {
	Name     string
	Center   bool
	Position templates.Position
	Size     templates.Size
}

func (i Image) Render(pdf *fpdf.Fpdf, pos RenderPosition) {
	x, y := Percent(pos.Width, float64(i.Position.X))+pos.X, Percent(pos.Height, float64(i.Position.Y))+pos.Y
	width, height := Percent(pos.Width, float64(i.Size.Width)), Percent(pos.Height, float64(i.Size.Height))

	info := pdf.GetImageInfo(i.Name)
	if i.Center && width == 0 {
		aspectRatio := float64(info.Width()) / float64(info.Height())
		width = height * aspectRatio
		x = x + (pos.Width-width)/2

		// prevent negative x, which would cause the image to be rendered on the opposite side
		x = math.Max(x, 0)
	} else if i.Center {
		x = x + (pos.Width-width)/2
	}

	pdf.ImageOptions(i.Name, x, y, width, height, false, fpdf.ImageOptions{}, 0, "")
}
