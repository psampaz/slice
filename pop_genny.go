package slice

import (
    "errors"
    "github.com/cheekybits/genny/generic"
)

type Type generic.Type

// PopType returns last value a Type slice and the remaining slice or an error in case of a nil or empty slice
func PopType(a []Type) (Type, []Type, error) {
    if len(a) == 0 {
        return 0, nil, errors.New("Cannot pop from a nil or empty slice")
    }
	
	return a[len(a)-1], a[:len(a)-1], nil
}