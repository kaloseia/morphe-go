package yaml

type EnumType string

const (
	EnumTypeString  EnumType = "String"
	EnumTypeInteger EnumType = "Integer"
	EnumTypeFloat   EnumType = "Float"
	EnumTypeBoolean EnumType = "Boolean"
	EnumTypeTime    EnumType = "Time"
	EnumTypeDate    EnumType = "Date"
)
