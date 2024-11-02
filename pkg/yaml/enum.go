package yaml

type Enum struct {
	Name    string         `yaml:"name"`
	Type    EnumType       `yaml:"type"`
	Entries map[string]any `yaml:"entries"`
}

func (e Enum) DeepClone() Enum {
	enumCopy := Enum{
		Name: e.Name,
		Type: e.Type,
	}

	entriesCopy := make(map[string]any, len(e.Entries))
	for key, primitive := range e.Entries {
		entriesCopy[key] = primitive
	}

	enumCopy.Entries = entriesCopy

	return enumCopy
}
