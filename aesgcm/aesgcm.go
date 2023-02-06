package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type Key struct {
	keyPhrase string
}

func NewKey(key string) *Key {
	return &Key{
		keyPhrase: md5Hashing(key),
	}
}

func (k *Key) Encrypt(plaintext, additional []byte) ([]byte, error) {
	return encrypt(k.keyPhrase, plaintext, additional)
}

func (k *Key) Decrypt(nonceCiphertext, additional []byte) ([]byte, error) {
	return decrypt(k.keyPhrase, nonceCiphertext, additional)
}

func md5Hashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:])
}

func encrypt(keyPhrase string, plaintext, additional []byte) ([]byte, error) {
	// The key phrase is hashed for increased security.
	block, err := aes.NewCipher([]byte(md5Hashing(keyPhrase)))
	if err != nil {
		return nil, err
	}

	// Returns a 128-bit block cipher with a nonce length.
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the plain text using the nonce.
	// Here, nonceCiphertext consist of both nonce + cipheredText, so that we don't have
	// to keep track of the nonce separately when decrypting.
	//
	// nonce := nonceCiphertext[:nonceSize]
	// ciphertext := nonceCiphertext[nonceSize:]
	nonceCiphertext := aesgcm.Seal(nonce, nonce, plaintext, additional)

	return nonceCiphertext, nil
}

func decrypt(keyPhrase string, nonceCiphertext, additional []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(md5Hashing(keyPhrase)))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := nonceCiphertext[:nonceSize], nonceCiphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, additional)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
