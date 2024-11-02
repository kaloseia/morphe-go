package yaml

import (
	"errors"
	"fmt"
)

var ErrNoMorpheModelName = errors.New("morphe model has no name")
var ErrNoMorpheModelFields = errors.New("morphe model has no fields")
var ErrNoMorpheModelIdentifiers = errors.New("morphe model has no identifiers")

func ErrMorpheModelUnknownFieldType(fieldName string, typeName string) error {
	return fmt.Errorf("morphe model field '%s' has unknown non-primitive type '%s'", fieldName, typeName)
}
