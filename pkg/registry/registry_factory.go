package registry

import "github.com/kaloseia/morphe-go/pkg/yaml"

func NewRegistry() *Registry {
	return &Registry{
		Models:   map[string]yaml.Model{},
		Entities: map[string]yaml.Entity{},
	}
}
