package yaml

import (
	"github.com/kaloseia/morphe-go/pkg/clone"
)

type Entity struct {
	Name    string                    `yaml:"name"`
	Fields  map[string]EntityField    `yaml:"fields"`
	Related map[string]EntityRelation `yaml:"related"`
}

func (e Entity) DeepClone() Entity {
	entityCopy := Entity{
		Name:    e.Name,
		Fields:  clone.DeepCloneMap(e.Fields),
		Related: clone.DeepCloneMap(e.Related),
	}

	return entityCopy
}
