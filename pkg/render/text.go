package render

import (
	"labl/pkg/templates"

	"github.com/go-pdf/fpdf"
)

func FitText(pdf *fpdf.Fpdf, text string, width float64, height float64) float64 {
	size := 0.5
	pdf.SetFontSize(size)

	for {
		if pdf.GetStringWidth(text) > width {
			break
		}
		_, lineHeight := pdf.GetFontSize()
		if lineHeight > height {
			break
		}

		size += 0.5
		pdf.SetFontSize(float64(size))
	}

	// should be at least 2 points smaller
	return float64(size - 2)
}

const DefaultFontName = "Arial"

type Text struct {
	Position templates.Position
	Size     templates.Size
	Font     templates.Font
	Color    [3]int
	Text     string
	Center   bool
}

func (t Text) Render(pdf *fpdf.Fpdf, pos RenderPosition) {
	if t.Font.Name == "" {
		t.Font.Name = DefaultFontName
	}

	pdf.SetFont(t.Font.Name, t.Font.Style, t.Font.Size)

	x, y := Percent(pos.Width, float64(t.Position.X))+pos.X, Percent(pos.Height, float64(t.Position.Y))+pos.Y
	width, height := Percent(pos.Width, float64(t.Size.Width)), Percent(pos.Height, float64(t.Size.Height))

	fontSize := t.Font.Size
	if fontSize == 0 {
		fontSize = FitText(pdf, t.Text, width, height)
	}

	align := "L"
	if t.Center {
		align = "C"
	}

	pdf.SetXY(x, y)
	pdf.SetFont(t.Font.Name, t.Font.Style, fontSize)
	pdf.SetTextColor(t.Color[0], t.Color[1], t.Color[2])

	pdf.CellFormat(width, height, t.Text, "", 0, align, false, 0, "")
}
