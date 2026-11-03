// Developer: Hamdy Abou El Anein
// Contact: hamdy.aea@protonmail.com
package main

import (
	"strings"
	"testing"
)

func TestPasswordLengthsAndAlphabet(t *testing.T) {
	for _, length := range []int{1, 2, 16, 128, 4096} {
		password, err := generatePassword(length)
		if err != nil || len(password) != length {
			t.Fatalf("length %d: got %d, error %v", length, len(password), err)
		}
		for _, character := range password {
			if !strings.ContainsRune(passwordAlphabet, character) {
				t.Fatalf("unexpected character %q", character)
			}
		}
	}
}

func TestInvalidLengths(t *testing.T) {
	for _, length := range []int{-1, 0, 4097} {
		if _, err := generatePassword(length); err == nil {
			t.Fatalf("accepted invalid length %d", length)
		}
	}
}
