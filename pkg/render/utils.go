package render

func Percent(value float64, percent float64) float64 {
	return value * percent / 100
}

type RenderPosition struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}
