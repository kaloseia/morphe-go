# morphe-go

Implementation wrapper to work with [Morphe specification](https://github.com/kaloseia/morphe) data in the Go programming language.

## Quick Start

1. Create your `.mod` and `.ent` YAML definition files in your project.
2. Install the `morphe-go` package via `go get github.com/kaloseia/morphe-go`.
3. Initialize the registry instance with `r := registry.NewRegistry()`.
4. Load models / entities with registry methods:
   * `modelsErr := r.LoadModelsFromDirectory(<modelsDirPath>)`
   * `entitiesErr := r.LoadEntitiesFromDirectory(<entitiesDirPath>)`
5. Access the registry model / entity definitions by name `<model|entity>, exists := r[<name>]` 
