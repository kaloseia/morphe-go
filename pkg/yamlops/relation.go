package yamlops

import "strings"

func IsRelationFor(relationType string) bool {
	return strings.Index(strings.ToLower(relationType), "for") == 0
}

func IsRelationHas(relationType string) bool {
	return strings.Index(strings.ToLower(relationType), "has") == 0
}

func IsRelationMany(relationType string) bool {
	return strings.Contains(strings.ToLower(relationType), "many")
}

func IsRelationOne(relationType string) bool {
	return strings.Contains(strings.ToLower(relationType), "one")
}
