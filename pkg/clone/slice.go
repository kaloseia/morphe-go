package clone

func Slice[TValue any](original []TValue) []TValue {
	fieldsCopy := make([]TValue, len(original))
	copy(fieldsCopy, original)
	return fieldsCopy
}
