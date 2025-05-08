package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"strings"
)

func ValidateUrl(rawURL string) error {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.New("invalid URL format")
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("invalid URL format")
	}
	return nil
}

func GenerateRandomCode(length int) (string, error) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := strings.Builder{}

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code.WriteByte(charset[num.Int64()])
	}
	return code.String(), nil
}

func Base62Encode(data []byte) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder
	num := new(big.Int).SetBytes(data)

	base := big.NewInt(62)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for num.Cmp(zero) > 0 {
		num.DivMod(num, base, mod)
		result.WriteByte(charset[mod.Int64()])
	}

	return result.String()
}
