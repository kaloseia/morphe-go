package yaml

import (
	"errors"
	"fmt"
)

var ErrNoMorpheEnumName = errors.New("morphe enum has no name")
var ErrNoMorpheEnumType = errors.New("morphe enum has no type")
var ErrNoMorpheEnumEntries = errors.New("morphe enum has no entries")

func ErrMorpheEnumUnsupportedEntryType(entryName string, entryValue any) error {
	return fmt.Errorf("enum entry '%s' value '%v' with type '%T' is not a primitive string or number type", entryName, entryValue, entryValue)
}

func ErrMorpheEnumEntryTypeMismatch(enumType EnumType, entryName string, entryValue any) error {
	return fmt.Errorf("enum entry '%s' value '%v' with type '%T' does not match the enum type of '%s'", entryName, entryValue, entryValue, enumType)
}
