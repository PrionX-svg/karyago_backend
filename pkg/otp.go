package pkg

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTPCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	minVal := int64(1)
	for i := 0; i < length-1; i++ {
		minVal *= 10
	}
	maxVal := minVal*10 - 1

	return strconv.FormatInt(r.Int63n(maxVal-minVal)+minVal, 10)
}
