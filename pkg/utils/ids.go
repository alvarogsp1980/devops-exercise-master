package utils

import (
	"fmt"
	"math/rand"
)

func NewID() string {
	return fmt.Sprintf("%v", rand.Int())
}
