package cfg

type MorpheLoadRegistryConfig struct {
	RegistryEnumsDirPath      string
	RegistryModelsDirPath     string
	RegistryStructuresDirPath string
	RegistryEntitiesDirPath   string
}

func (config MorpheLoadRegistryConfig) Validate() error {
	if config.RegistryEnumsDirPath == "" {
		return ErrNoRegistryEnumsDirPath
	}
	if config.RegistryModelsDirPath == "" {
		return ErrNoRegistryModelsDirPath
	}
	if config.RegistryStructuresDirPath == "" {
		return ErrNoRegistryStructuresDirPath
	}
	if config.RegistryEntitiesDirPath == "" {
		return ErrNoRegistryEntitiesDirPath
	}
	return nil
}
