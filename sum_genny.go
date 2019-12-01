package slice

import (
	"errors"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

// SumType returns the sum of the values of a Type Slice or an error in case of a nil or empty slice
func SumType(a []Type) (Type, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot calculate the sum of a nil or empty slice")
	}

	var sum Type
	for k := range a {
		sum += a[k]
	}

	return sum, nil
}
