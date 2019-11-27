package slice

import "github.com/cheekybits/genny/generic"

type Type generic.Type

// ReverseType reverses a Type slice
func ReverseType(a []Type) []Type {
	if len(a) == 0 {
		return a
	}

	s, e := 0, len(a)-1
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}

	return a
}