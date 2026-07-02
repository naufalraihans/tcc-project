package utils

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const orderChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func OrderID() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = orderChars[rand.IntN(len(orderChars))]
	}
	return fmt.Sprintf("TCC-%s-%d", b, time.Now().Unix())
}
