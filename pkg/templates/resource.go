package templates

type Source struct {
	Type  SourceType `json:"type"`
	Value string     `json:"value"`
}

type Resource struct {
	Type   ResourceType `json:"type"`
	Label  string       `json:"label"`
	Source Source       `json:"source"`
}
