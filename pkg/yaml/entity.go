package yaml

type Entity struct {
	Name    string                    `yaml:"name"`
	Fields  map[string]EntityField    `yaml:"fields"`
	Related map[string]EntityRelation `yaml:"related"`
}
