package common

import (
	"math/rand"
)

func CreateId() int64{
	return rand.Int63()
}
