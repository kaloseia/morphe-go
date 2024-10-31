package yaml

type EnumEntry map[string]any

func (e EnumEntry) DeepClone() EnumEntry {
	entryCopy := make(map[string]any, len(e))
	for key, primitive := range e {
		entryCopy[key] = primitive
	}
	return entryCopy
}
