// Developer: Hamdy Abou El Anein
// Contact: hamdy.aea@protonmail.com
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const passwordAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$;?0123456789"

func generatePassword(n int) (string, error) {
	if n < 1 || n > 4096 {
		return "", fmt.Errorf("password length must be between 1 and 4096")
	}
	password := make([]byte, n)
	limit := big.NewInt(int64(len(passwordAlphabet)))
	for i := range password {
		index, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return "", err
		}
		password[i] = passwordAlphabet[index.Int64()]
	}
	return string(password), nil
}

func main() {
	password, err := generatePassword(16)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(password)
}
