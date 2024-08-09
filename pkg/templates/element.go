package templates

type ElementOptions struct {
	Center           bool   `json:"center"`
	CenterVertical   bool   `json:"centerVertical"`
	CenterHorizontal bool   `json:"centerHorizontal"`
	Font             Font   `json:"font"`
	Color            [3]int `json:"color"`
}

type Element struct {
	Type ElementType `json:"type"`
	// The name of the resource
	Resource      string         `json:"resource"`
	ColorResource string         `json:"colorResource"`
	Position      Position       `json:"position"`
	Size          Size           `json:"size"`
	Options       ElementOptions `json:"options"`
}
