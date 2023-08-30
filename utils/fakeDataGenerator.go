package utils

import (
	"math"
	"math/rand"
	"time"
)

// Generates future date (but before 2022)
func GenerateRandomFutureDate() time.Time {
	min := time.Now().Unix()
	max := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenerateRandomFee() float64 {
	x := rand.Float64() * 100
	return math.Ceil(x*100) / 100

}
