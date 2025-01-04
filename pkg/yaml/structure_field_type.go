package yaml

import "slices"

type StructureFieldType string

const (
	StructureFieldTypeUUID          StructureFieldType = "UUID"
	StructureFieldTypeAutoIncrement StructureFieldType = "AutoIncrement"
	StructureFieldTypeString        StructureFieldType = "String"
	StructureFieldTypeInteger       StructureFieldType = "Integer"
	StructureFieldTypeFloat         StructureFieldType = "Float"
	StructureFieldTypeBoolean       StructureFieldType = "Boolean"
	StructureFieldTypeTime          StructureFieldType = "Time"
	StructureFieldTypeDate          StructureFieldType = "Date"
	StructureFieldTypeProtected     StructureFieldType = "Protected"
	StructureFieldTypeSealed        StructureFieldType = "Sealed"
)

var StructureFieldTypesPrimitive = []StructureFieldType{
	StructureFieldTypeUUID,
	StructureFieldTypeAutoIncrement,
	StructureFieldTypeString,
	StructureFieldTypeInteger,
	StructureFieldTypeFloat,
	StructureFieldTypeBoolean,
	StructureFieldTypeTime,
	StructureFieldTypeDate,
	StructureFieldTypeProtected,
	StructureFieldTypeSealed,
}

func IsStructureFieldTypePrimitive(t StructureFieldType) bool {
	return slices.Contains(StructureFieldTypesPrimitive, t)
}
