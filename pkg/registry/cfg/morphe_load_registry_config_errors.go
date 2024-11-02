package cfg

import "errors"

var ErrNoRegistryEnumsDirPath = errors.New("registry enums dir path cannot be empty")
var ErrNoRegistryModelsDirPath = errors.New("registry models dir path cannot be empty")
var ErrNoRegistryEntitiesDirPath = errors.New("registry entities dir cannot be empty")
