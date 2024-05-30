package clone

func DeepCloneMap[TKey comparable, TCloneable DeepCloneable[TCloneable]](original map[TKey]TCloneable) map[TKey]TCloneable {
	copiedMap := make(map[TKey]TCloneable, len(original))
	for key, copyable := range original {
		copiedMap[key] = copyable.DeepClone()
	}
	return copiedMap
}
