package slice

import "github.com/cheekybits/genny/generic"

type Type generic.Type

// FilterType
func FilterType(a []Type, keep func(x Type) bool) []Type {
	if len(a) == 0 {
		return a
	}

	n := 0
	for _, v := range a {
		if keep(v) {
			a[n] = v
			n++
		}
	}

	return a[:n]
}