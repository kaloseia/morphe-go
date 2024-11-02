package yaml

import "slices"

type ModelFieldType string

const (
	ModelFieldTypeUUID          ModelFieldType = "UUID"
	ModelFieldTypeAutoIncrement ModelFieldType = "AutoIncrement"
	ModelFieldTypeString        ModelFieldType = "String"
	ModelFieldTypeInteger       ModelFieldType = "Integer"
	ModelFieldTypeFloat         ModelFieldType = "Float"
	ModelFieldTypeBoolean       ModelFieldType = "Boolean"
	ModelFieldTypeTime          ModelFieldType = "Time"
	ModelFieldTypeDate          ModelFieldType = "Date"
	ModelFieldTypeProtected     ModelFieldType = "Protected"
	ModelFieldTypeSealed        ModelFieldType = "Sealed"
)

var ModelFieldTypesPrimitive = []ModelFieldType{
	ModelFieldTypeUUID,
	ModelFieldTypeAutoIncrement,
	ModelFieldTypeString,
	ModelFieldTypeInteger,
	ModelFieldTypeFloat,
	ModelFieldTypeBoolean,
	ModelFieldTypeTime,
	ModelFieldTypeDate,
	ModelFieldTypeProtected,
	ModelFieldTypeSealed,
}

func IsModelFieldTypePrimitive(t ModelFieldType) bool {
	return slices.Contains(ModelFieldTypesPrimitive, t)
}
