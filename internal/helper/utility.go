package helper

import (
	"crypto/rand"
	"math/big"
)

func RandomNumbers(length int) (int, error) {
	const numbers = "1234567890"
	if length <= 0 {
		return 0, nil
	}

	// Generate a random number with the specified number of digits
	max := big.NewInt(1)
	for i := 0; i < length; i++ {
		max.Mul(max, big.NewInt(10))
	}
	max.Sub(max, big.NewInt(1)) // Subtract 1 to get the maximum n-digit number

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	// Ensure the number has the exact number of digits by padding with leading zeros if needed
	min := big.NewInt(1)
	for i := 1; i < length; i++ {
		min.Mul(min, big.NewInt(10))
	}
	if n.Cmp(min) < 0 {
		n.Add(n, min)
	}

	return int(n.Int64()), nil
}
