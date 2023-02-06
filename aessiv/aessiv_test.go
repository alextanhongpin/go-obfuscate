package aessiv_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/alextanhongpin/go-obfuscate/aessiv"
)

func TestEncryptDecrypt(t *testing.T) {
	key := aessiv.NewKey(strings.Repeat("A", 32))
	data := []byte("hello world")
	additional := []byte("v1")

	enc, err := key.Encrypt(data, additional)
	if err != nil {
		t.Error(err)
	}

	dec, err := key.Decrypt(enc, additional)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, dec) {
		t.Fatalf("expected %s, got %s", data, dec)
	}
}
