package yaml

import "errors"

var ErrNoMorpheModelName = errors.New("morphe model has no name")
var ErrNoMorpheModelFields = errors.New("morphe model has no fields")
var ErrNoMorpheModelIdentifiers = errors.New("morphe model has no identifiers")
