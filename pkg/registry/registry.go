package registry

import (
	"fmt"
	"sync"

	"github.com/kaloseia/morphe-go/pkg/clone"
	"github.com/kaloseia/morphe-go/pkg/yaml"
	"github.com/kaloseia/morphe-go/pkg/yamlfile"
)

const ModelFileSuffix = ".mod"
const EntityFileSuffix = ".ent"

type Registry struct {
	mutex sync.RWMutex

	models   map[string]yaml.Model  `yaml:"models"`
	entities map[string]yaml.Entity `yaml:"entities"`
}

// SetModel is a thread-safe way to write a model to the registry
func (r *Registry) SetModel(name string, model yaml.Model) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.models[name] = model
}

// GetModel returns a thread-safe copy of a registry model
func (r *Registry) GetModel(name string) (yaml.Model, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	model, modelFound := r.models[name]
	if !modelFound {
		return yaml.Model{}, fmt.Errorf("model with name '%s' not found registry", name)
	}
	modelClone := model.DeepClone()
	return modelClone, nil
}

// GetAllModels returns a thread-safe copy of all registry models
func (r *Registry) GetAllModels() map[string]yaml.Model {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	modelsClone := clone.DeepCloneMap(r.models)
	return modelsClone
}

// SetEntity is a thread-safe way to write an entity to the registry
func (r *Registry) SetEntity(name string, entity yaml.Entity) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.entities[name] = entity
}

// GetEntity returns a thread-safe copy of a registry entity
func (r *Registry) GetEntity(name string) (yaml.Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	entity, entityFound := r.entities[name]
	if !entityFound {
		return yaml.Entity{}, fmt.Errorf("entity with name '%s' not found registry", name)
	}
	entityClone := entity.DeepClone()
	return entityClone, nil
}

// GetAllEntities returns a thread-safe copy of all registry entities
func (r *Registry) GetAllEntities() map[string]yaml.Entity {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	entitiesClone := clone.DeepCloneMap(r.entities)
	return entitiesClone
}

func (r *Registry) DeepClone() *Registry {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	registryCopy := &Registry{
		models:   clone.DeepCloneMap(r.models),
		entities: clone.DeepCloneMap(r.entities),
	}

	return registryCopy
}

func (r *Registry) LoadModelsFromDirectory(dirPath string) error {
	allModels, unmarshalErr := yamlfile.UnmarshalAllYAMLFiles[yaml.Model](dirPath, ModelFileSuffix)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	loadErr := r.loadModelDefinitions(allModels)
	return loadErr
}

func (r *Registry) LoadEntitiesFromDirectory(dirPath string) error {
	allEntities, unmarshalErr := yamlfile.UnmarshalAllYAMLFiles[yaml.Entity](dirPath, EntityFileSuffix)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	loadErr := r.loadEntityDefinitions(allEntities)
	return loadErr
}

func (r *Registry) loadModelDefinitions(allModels map[string]yaml.Model) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.models == nil {
		r.models = make(map[string]yaml.Model)
	}

	for modelPathAbs, model := range allModels {
		_, nameConflict := r.models[model.Name]
		if nameConflict {
			return fmt.Errorf("model name '%s' already exists in registry (conflict: %s)", model.Name, modelPathAbs)
		}

		r.models[model.Name] = model
	}
	return nil
}

func (r *Registry) loadEntityDefinitions(allEntities map[string]yaml.Entity) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.entities == nil {
		r.entities = make(map[string]yaml.Entity)
	}

	for entityPathAbs, entity := range allEntities {
		_, nameConflict := r.entities[entity.Name]
		if nameConflict {
			return fmt.Errorf("entity name '%s' already exists in registry (conflict: %s)", entity.Name, entityPathAbs)
		}

		r.entities[entity.Name] = entity
	}
	return nil
}
