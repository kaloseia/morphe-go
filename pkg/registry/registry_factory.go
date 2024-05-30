package registry

import "github.com/kaloseia/morphe-go/pkg/yaml"

func NewRegistry() *Registry {
	return &Registry{
		models:   map[string]yaml.Model{},
		entities: map[string]yaml.Entity{},
	}
}
