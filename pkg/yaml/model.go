package yaml

import (
	"log"

	"github.com/kaloseia/clone"
)

type Model struct {
	Name        string                     `yaml:"name"`
	Fields      map[string]ModelField      `yaml:"fields"`
	Identifiers map[string]ModelIdentifier `yaml:"identifiers"`
	Related     map[string]ModelRelation   `yaml:"related"`
}

func (m Model) Validate() error {
	if m.Name == "" {
		return ErrNoMorpheModelName
	}
	if len(m.Fields) == 0 {
		return ErrNoMorpheModelFields
	}
	if len(m.Identifiers) == 0 {
		return ErrNoMorpheModelIdentifiers
	}
	return nil
}

func (m Model) DeepClone() Model {
	modelCopy := Model{
		Name:        m.Name,
		Fields:      clone.DeepCloneMap(m.Fields),
		Identifiers: clone.DeepCloneMap(m.Identifiers),
		Related:     clone.DeepCloneMap(m.Related),
	}

	return modelCopy
}

func (m Model) GetIdentifierFields() []ModelField {
	var fields []ModelField
	for _, identifier := range m.Identifiers {
		for _, fieldName := range identifier.Fields {
			idField, fieldExists := m.Fields[fieldName]
			if !fieldExists {
				log.Printf("identifier field '%s' does not exist in model '%s'", fieldName, m.Name)
				continue
			}
			fields = append(fields, idField)
		}
	}
	return fields
}
