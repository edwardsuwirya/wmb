package utils

import (
	"fmt"
	"math/rand"
)

func GenerateId() string {
	return fmt.Sprintf("%d", rand.Int())
}
