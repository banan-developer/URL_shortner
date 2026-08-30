package generation

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerationCode(lenght int) (string, error) {
	result := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}

		result[i] = alphabet[n.Int64()]
	}
	return string(result), nil
}
