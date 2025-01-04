package yaml

import (
	"errors"
	"fmt"
)

var ErrNoMorpheStructureName = errors.New("morphe structure has no name")
var ErrNoMorpheStructureFields = errors.New("morphe structure has no fields")

func ErrMorpheStructureUnknownFieldType(fieldName string, typeName string) error {
	return fmt.Errorf("morphe structure field '%s' has unknown non-primitive type '%s'", fieldName, typeName)
}
