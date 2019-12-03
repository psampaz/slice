package slice

import (
	"math/rand"
	"time"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

// ShuffleType shuffles (in place) a Type slice
func ShuffleType(a []Type) []Type {
	if len(a) <= 1 {
		return a
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}
