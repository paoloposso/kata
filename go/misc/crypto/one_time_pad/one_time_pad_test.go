package onetimepad

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	encrypted, dummy := Encrypt("Paolo")

	fmt.Println(encrypted)
	fmt.Println(dummy)
}
