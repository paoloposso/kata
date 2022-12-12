package onetimepad

import (
	"crypto/rand"
)

func getRandomKey(length int) []byte {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return bytes
}

func Encrypt(original string) ([]byte, []byte) {
	randomKey := getRandomKey(len(original))
	originalKey := []byte(original)

	encrypted := make([]byte, len(original))
	for i := range encrypted {
		encrypted[i] = originalKey[i] ^ randomKey[i]
	}

	return randomKey, encrypted
}

func Decrypt(randomKey []byte, encrypted []byte) string {
	decrypted := make([]byte, len(encrypted))
	for i := range decrypted {
		decrypted[i] = encrypted[i] ^ randomKey[i]
	}
	return string(decrypted[:])
}
