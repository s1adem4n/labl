package render

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"labl/pkg/templates"

	"github.com/go-pdf/fpdf"
)

type Renderer interface {
	Render(pdf *fpdf.Fpdf, pos RenderPosition)
}

type Size struct {
	Width  float64
	Height float64
}

type TemplateRenderer struct {
	PDF       *fpdf.Fpdf
	Inputs    Inputs
	Template  *templates.Template
	Resources Resources
	Gap       float64
	Margin    float64
	Quantity  int
	Size      Size
	Outline   bool
}

func (t *TemplateRenderer) ValidateAspectRatio() error {
	if t.Size.Height == 0 && t.Size.Width == 0 {
		return fmt.Errorf("invalid size: %f x %f", t.Size.Width, t.Size.Height)
	}

	aspectRatio := t.Template.AspectRatio
	if aspectRatio[0] == 0 || aspectRatio[1] == 0 {
		return fmt.Errorf("invalid aspect ratio: %d:%d", aspectRatio[0], aspectRatio[1])
	}
	w, l := float64(aspectRatio[0]), float64(aspectRatio[1])

	expectedWidth := t.Size.Height * (w / l)
	expectedHeight := t.Size.Width * (l / w)

	if t.Size.Width == 0 {
		t.Size.Width = expectedWidth
		return nil
	} else if t.Size.Height == 0 {
		t.Size.Height = expectedHeight
		return nil
	}

	if t.Size.Height != expectedHeight || t.Size.Width != expectedWidth {
		return fmt.Errorf("invalid aspect ratio, expected %d:%d", aspectRatio[0], aspectRatio[1])
	}

	return nil
}

func (t *TemplateRenderer) LoadFonts() error {
	for name, r := range t.Template.Resources {
		if r.Type != templates.ResourceTypeFont {
			continue
		}

		v, err := t.Resources.GetBytes(name)
		if err != nil {
			return err
		}
		t.PDF.AddUTF8FontFromBytes(name, "", v)
	}
	return nil
}

func GetImageType(data []byte) (string, error) {
	_, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return format, nil
}

func (t *TemplateRenderer) LoadImages() error {
	for name, r := range t.Template.Resources {
		if r.Type != templates.ResourceTypeImage {
			continue
		}

		v, err := t.Resources.GetBytes(name)
		if err != nil {
			return err
		}
		format, err := GetImageType(v)
		if err != nil {
			return err
		}

		reader := bytes.NewReader(v)
		t.PDF.RegisterImageOptionsReader(name, fpdf.ImageOptions{
			ImageType: format,
			ReadDpi:   true,
		}, reader)
	}
	return nil
}

func (t *TemplateRenderer) Render() error {
	if err := t.ValidateAspectRatio(); err != nil {
		return err
	}
	if err := t.Resources.AddResources(t.Template.Resources, t.Inputs); err != nil {
		return err
	}
	if err := t.LoadFonts(); err != nil {
		return err
	}
	if err := t.LoadImages(); err != nil {
		return err
	}

	pdf := t.PDF
	pdf.AddPage()

	width, height := pdf.GetPageSize()
	x, y := t.Margin, t.Margin

	for i := 0; i < t.Quantity; i++ {
		if x+t.Size.Width+t.Gap+t.Margin > width {
			x = t.Margin
			y += t.Size.Height + t.Gap
		}
		if y+t.Size.Height+t.Gap+t.Margin > height {
			pdf.AddPage()
			x = t.Margin
			y = t.Margin
		}

		pos := RenderPosition{
			X:      x,
			Y:      y,
			Width:  t.Size.Width,
			Height: t.Size.Height,
		}

		for _, e := range t.Template.Elements {
			switch e.Type {
			case templates.ElementTypeText:
				v, err := t.Resources.GetString(e.Resource)
				if err != nil {
					return err
				}
				text := Text{
					Position: e.Position,
					Size:     e.Size,
					Font:     e.Options.Font,
					Color:    e.Options.Color,
					Text:     v,
					Center:   e.Options.Center,
				}
				text.Render(pdf, pos)
			case templates.ElementTypeImage:
				image := Image{
					Name:             e.Resource,
					Position:         e.Position,
					CenterVertical:   e.Options.CenterVertical,
					CenterHorizontal: e.Options.CenterHorizontal,
					Center:           e.Options.Center,
					Size:             e.Size,
				}
				image.Render(pdf, pos)
			}
		}

		if t.Outline {
			pdf.Rect(x, y, t.Size.Width, t.Size.Height, "D")
		}

		x += t.Size.Width + t.Gap
	}

	return nil
}
