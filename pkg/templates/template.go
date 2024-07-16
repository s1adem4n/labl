package templates

type Template struct {
	AspectRatio [2]int              `json:"aspectRatio"`
	Resources   map[string]Resource `json:"resources"`
	Elements    []Element           `json:"elements"`
}
