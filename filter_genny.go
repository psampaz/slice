package slice

import "github.com/cheekybits/genny/generic"

type Type generic.Type

// FilterType
func FilterType(a []Type, keep func(x Type) bool) {
	n := 0
	for k := range a {
		if keep(a[k]) {
			a[n] = a[k]
			n++
		}
	}

	return a[:n]
}
