package utility

import (
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}

func GenerateOrderCode() string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digitBytes = "0123456789"

	rand.Seed(time.Now().UnixNano())

	// Generate 8 random digits
	digits := make([]byte, 8)
	for i := range digits {
		digits[i] = digitBytes[rand.Intn(len(digitBytes))]
	}

	// Generate 4 random letters
	letters := make([]byte, 4)
	for i := range letters {
		letters[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	// Combine parts into the final string
	var sb strings.Builder
	sb.WriteString("KSH")
	sb.Write(digits)
	sb.Write(letters)

	return sb.String()
}
