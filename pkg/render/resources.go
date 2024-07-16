package render

import (
	"fmt"
	"labl/pkg/assets"
	"labl/pkg/templates"
)

func ValidateResource(resource templates.Resource, value any) bool {
	switch resource.Type {
	case templates.ResourceTypeText:
		if _, ok := value.(string); !ok {
			return false
		}
	case templates.ResourceTypeImage:
		if _, ok := value.([]byte); !ok {
			return false
		}
	case templates.ResourceTypeFont:
		if _, ok := value.([]byte); !ok {
			return false
		}
	}
	return true
}

type Inputs map[string]any
type Resources map[string]any

func (r Resources) AddResource(name string, resource templates.Resource, inputs Inputs) error {
	var valid bool
	var value any = nil

	switch resource.Source.Type {
	case templates.SourceTypeEmbed:
		if v, ok := assets.Assets[resource.Source.Value]; ok {
			value = v
		}
	case templates.SourceTypeInput:
		if v, ok := inputs[name]; ok {
			value = v
		}
	case templates.SourceTypeCustom:
		if v, ok := inputs[resource.Source.Value]; ok {
			value = v
		}
	}

	if value == nil {
		return fmt.Errorf("missing resource: %s", resource.Source.Value)
	}
	if valid = ValidateResource(resource, value); !valid {
		return fmt.Errorf("invalid value for resource %s", resource.Source.Value)
	}
	if _, ok := r[name]; !ok {
		r[name] = value
	} else {
		return fmt.Errorf("resource already exists: %s", name)
	}

	return nil
}

func (r Resources) AddResources(resources map[string]templates.Resource, inputs Inputs) error {
	for name, resource := range resources {
		err := r.AddResource(name, resource, inputs)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r Resources) GetString(k string) (string, error) {
	if v, ok := r[k].(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("resource %s is not a string", k)
}

func (r Resources) GetBytes(k string) ([]byte, error) {
	if v, ok := r[k].([]byte); ok {
		return v, nil
	}
	return nil, fmt.Errorf("resource %s is not a byte slice", k)
}
