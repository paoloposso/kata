package onetimepad

import (
	"crypto/rand"
	"encoding/binary"
)

func getRandomKey(length int) uint32 {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return binary.BigEndian.Uint32(bytes)
}

func Encrypt(original string) (int32, int32) {
	dummy := getRandomKey(len(original))
	originalKey := binary.BigEndian.Uint32([]byte(original))
	encrypted := originalKey ^ dummy
	return int32(encrypted), int32(dummy)
}
