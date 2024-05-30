package yaml

type ModelRelation struct {
	Type string `yaml:"type"`
}

func (r ModelRelation) DeepClone() ModelRelation {
	return ModelRelation{
		Type: r.Type,
	}
}
