package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateID creates a unique 5-digit string ID.
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%05d", rand.Intn(100000))
}
