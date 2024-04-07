package registry

import (
	"fmt"

	"github.com/kaloseia/morphe-go/pkg/yaml"
	"github.com/kaloseia/morphe-go/pkg/yamlfile"
)

const ModelFileSuffix = ".mod"
const EntityFileSuffix = ".ent"

type Registry struct {
	Models   map[string]yaml.Model  `yaml:"models"`
	Entities map[string]yaml.Entity `yaml:"entities"`
}

func (registry *Registry) LoadModelsFromDirectory(dirPath string) error {
	allModels, unmarshalErr := yamlfile.UnmarshalAllYAMLFiles[yaml.Model](dirPath, ModelFileSuffix)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	loadErr := registry.loadModelDefinitions(allModels)
	return loadErr
}

func (registry *Registry) LoadEntitiesFromDirectory(dirPath string) error {
	allEntities, unmarshalErr := yamlfile.UnmarshalAllYAMLFiles[yaml.Entity](dirPath, EntityFileSuffix)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	loadErr := registry.loadEntityDefinitions(allEntities)
	return loadErr
}

func (registry *Registry) loadModelDefinitions(allModels map[string]yaml.Model) error {
	if registry.Models == nil {
		registry.Models = make(map[string]yaml.Model)
	}

	for modelPathAbs, model := range allModels {
		_, nameConflict := registry.Models[model.Name]
		if nameConflict {
			return fmt.Errorf("model name '%s' already exists in registry (conflict: %s)", model.Name, modelPathAbs)
		}

		registry.Models[model.Name] = model
	}
	return nil
}

func (registry *Registry) loadEntityDefinitions(allEntities map[string]yaml.Entity) error {
	if registry.Entities == nil {
		registry.Entities = make(map[string]yaml.Entity)
	}

	for entityPathAbs, entity := range allEntities {
		_, nameConflict := registry.Entities[entity.Name]
		if nameConflict {
			return fmt.Errorf("entity name '%s' already exists in registry (conflict: %s)", entity.Name, entityPathAbs)
		}

		registry.Entities[entity.Name] = entity
	}
	return nil
}
