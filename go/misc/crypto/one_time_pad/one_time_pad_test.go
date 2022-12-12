package onetimepad

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	original := "PAOLO"

	encrypted, dummy := Encrypt(original)
	dec := Decrypt(dummy, encrypted)

	if original != dec {
		t.Fail()
	}
}
