package yaml

type EntityField struct {
	Type       ModelFieldPath `yaml:"type"`
	Attributes []string       `yaml:"attributes"`
}
