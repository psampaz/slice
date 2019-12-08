package slice

import (
	"errors"
	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

// DeleteRangeType deletes the elements between from and to index (inclusive) from a Type slice
func DeleteRangeType(a []Type, from, to Type) ([]Type, error) {
	if len(a) == 0 {
		return nil, errors.New("Can't cut a nil or empty slice")
	}

	if from < 0 {
		return nil, errors.New("from index out of bounds")
	}

	if to > len(a)-1 {
		return nil, errors.New("to index out of bounds")
	}

	if from > to {
		return nil, errors.New("from index must be less or equal to to index")
	}

	return append(a[:from], a[to+1:]...), nil
}
