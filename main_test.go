package main

import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"strings"
	"testing"

	"github.com/alextanhongpin/go-obfuscate/aesgcm"
	"github.com/alextanhongpin/go-obfuscate/aessiv"
	"github.com/alextanhongpin/go-obfuscate/hashid"
)

func BenchmarkHashID(b *testing.B) {
	key := hashid.NewKey(strings.Repeat("A", 32), 6)

	for i := 0; i < b.N; i++ {
		p := mathrand.Int63()
		n := mathrand.Int63()
		enc, err := key.Encode(p, n)
		if err != nil {
			b.Error(err)
		}

		dec, err := key.Decode(p, enc)
		if err != nil {
			b.Error(err)
		}

		if n != dec {
			b.Fatalf("expected %d, got %d", n, dec)
		}
	}
}

func BenchmarkDaead(b *testing.B) {
	key := aessiv.NewKey(strings.Repeat("A", 64))

	for i := 0; i < b.N; i++ {
		n := mathrand.Int63()
		data := []byte(fmt.Sprint(n))
		ciphertext, err := key.Encrypt(data, nil)
		if err != nil {
			b.Error(err)
		}

		decrypted, err := key.Decrypt(ciphertext, nil)
		if err != nil {
			b.Error(err)
		}
		if !bytes.Equal(data, decrypted) {
			b.Fatalf("expected %s, got %s", data, decrypted)
		}
	}
}

func BenchmarkAESGCM(b *testing.B) {
	key := aesgcm.NewKey(strings.Repeat("A", 32))

	for i := 0; i < b.N; i++ {
		n := mathrand.Int63()
		data := []byte(fmt.Sprint(n))
		ciphertext, err := key.Encrypt(data, nil)
		if err != nil {
			b.Error(err)
		}

		decrypted, err := key.Decrypt(ciphertext, nil)
		if err != nil {
			b.Error(err)
		}
		if !bytes.Equal(data, decrypted) {
			b.Fatalf("expected %s, got %s", data, decrypted)
		}
	}
}
