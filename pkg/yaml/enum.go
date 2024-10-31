package yaml

import "github.com/kaloseia/clone"

type Enum struct {
	Name    string               `yaml:"name"`
	Type    EnumType             `yaml:"type"`
	Entries map[string]EnumEntry `yaml:"entries"`
}

func (e Enum) DeepClone() Enum {
	enumCopy := Enum{
		Name:    e.Name,
		Entries: clone.DeepCloneMap(e.Entries),
	}

	return enumCopy
}
