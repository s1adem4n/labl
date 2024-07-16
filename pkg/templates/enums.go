package templates

import "fmt"

type ElementType int

const (
	ElementTypeText ElementType = iota
	ElementTypeImage
)

var ElementTypeStrings = map[ElementType]string{
	ElementTypeText:  "text",
	ElementTypeImage: "image",
}
var ElementTypeValues = map[string]ElementType{
	"text":  ElementTypeText,
	"image": ElementTypeImage,
}

func (r ElementType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ElementTypeStrings[r] + `"`), nil
}
func (r *ElementType) UnmarshalJSON(b []byte) error {
	// remove quotes
	b = b[1 : len(b)-1]

	for k, v := range ElementTypeValues {
		if k == string(b) {
			*r = v
			return nil
		}
	}
	return fmt.Errorf("unknown ElementType: %s", string(b))
}
func (r ElementType) String() string {
	return ElementTypeStrings[r]
}

type ResourceType int

const (
	ResourceTypeText ResourceType = iota
	ResourceTypeImage
	ResourceTypeFont
)

var ResourceTypeStrings = map[ResourceType]string{
	ResourceTypeText:  "text",
	ResourceTypeImage: "image",
	ResourceTypeFont:  "font",
}
var ResourceTypeValues = map[string]ResourceType{
	"text":  ResourceTypeText,
	"image": ResourceTypeImage,
	"font":  ResourceTypeFont,
}

func (r ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ResourceTypeStrings[r] + `"`), nil
}
func (r *ResourceType) UnmarshalJSON(b []byte) error {
	// remove quotes
	b = b[1 : len(b)-1]

	for k, v := range ResourceTypeValues {
		if k == string(b) {
			*r = v
			return nil
		}
	}
	return fmt.Errorf("unknown ResourceType: %s", string(b))
}
func (r ResourceType) String() string {
	return ResourceTypeStrings[r]
}

type SourceType int

const (
	SourceTypeEmbed SourceType = iota
	SourceTypeInput
	SourceTypeCustom
)

var SourceTypeStrings = map[SourceType]string{
	SourceTypeEmbed:  "embed",
	SourceTypeInput:  "input",
	SourceTypeCustom: "custom",
}
var SourceTypeValues = map[string]SourceType{
	"embed":  SourceTypeEmbed,
	"input":  SourceTypeInput,
	"custom": SourceTypeCustom,
}

func (s SourceType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + SourceTypeStrings[s] + `"`), nil
}
func (s *SourceType) UnmarshalJSON(b []byte) error {
	// remove quotes
	b = b[1 : len(b)-1]

	for k, v := range SourceTypeValues {
		if k == string(b) {
			*s = v
			return nil
		}
	}

	return fmt.Errorf("unknown SourceType: %s", string(b))
}
func (s SourceType) String() string {
	return SourceTypeStrings[s]
}
