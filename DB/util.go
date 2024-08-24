package DB

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() string {
	timestamp := time.Now().UnixNano()
	random := rand.Int63()
	return fmt.Sprintf("%x-%x", timestamp, random)
}
