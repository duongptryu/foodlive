package common

import "math/rand"

const (
	EntityOTP = "OTP"
	numberStr = "0123456789"
)

func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numberStr[rand.Intn(len(numberStr))]
	}
	return string(b)
}
