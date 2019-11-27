package slice

import "github.com/cheekybits/genny/generic"

type Type generic.Type

// DeduplicateType deduplicates (in place) a Type slice preserving the original order 
func DeduplicateType(a []Type) []Type {
	if len(a) == 0 {
		return a
	}

	seen := make(map[Type]struct{}, 0)
	
	j := 0
	for k := range a {
		if _, ok := seen[a[k]]; ok {
			continue
		}
		seen[a[k]] = struct{}{}
		a[j] = a[k]
		j++
	}

	return a[:j]
}