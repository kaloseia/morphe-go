package cfg

type MorpheLoadRegistryConfig struct {
	RegistryModelsDirPath   string
	RegistryEntitiesDirPath string
}

func (config MorpheLoadRegistryConfig) Validate() error {
	if config.RegistryModelsDirPath == "" {
		return ErrNoRegistryModelsDirPath
	}
	if config.RegistryEntitiesDirPath == "" {
		return ErrNoRegistryEntitiesDirPath
	}
	return nil
}
