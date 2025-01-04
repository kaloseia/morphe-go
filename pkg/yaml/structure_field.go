package yaml

import "github.com/kaloseia/clone"

type StructureField struct {
	Type       StructureFieldType `yaml:"type"`
	Attributes []string           `yaml:"attributes"`
}

func (f StructureField) DeepClone() StructureField {
	return StructureField{
		Type:       f.Type,
		Attributes: clone.Slice(f.Attributes),
	}
}
