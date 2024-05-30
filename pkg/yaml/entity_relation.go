package yaml

type EntityRelation struct {
	Type string `yaml:"type"`
}

func (f EntityRelation) DeepClone() EntityRelation {
	return EntityRelation{
		Type: f.Type,
	}
}
