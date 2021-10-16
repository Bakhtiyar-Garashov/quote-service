package utils

import (
	"math/rand"
	"time"
)

// Generates date between now and 2030
func GenerateRandomFutureDate() time.Time {
	min := time.Now().Unix()
	max := time.Date(2030, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenerateRandomFee() float64 {
	return rand.Float64() * 100
}
