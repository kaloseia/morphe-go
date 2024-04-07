package yaml

type ModelField struct {
	Type       ModelFieldType `yaml:"type"`
	Attributes []string       `yaml:"attributes"`
}
