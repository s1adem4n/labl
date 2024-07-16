package templates

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Font struct {
	Name  string  `json:"name"`
	Style string  `json:"style"`
	Size  float64 `json:"size"`
}
