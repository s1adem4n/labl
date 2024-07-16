package templates

type ElementOptions struct {
	Center bool   `json:"center"`
	Font   Font   `json:"font"`
	Color  [3]int `json:"color"`
}

type Element struct {
	Type ElementType `json:"type"`
	// The name of the resource
	Resource string         `json:"resource"`
	Position Position       `json:"position"`
	Size     Size           `json:"size"`
	Options  ElementOptions `json:"options"`
}
