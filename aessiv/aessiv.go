package aessiv

import "github.com/google/tink/go/daead/subtle"

type Key struct {
	kh *subtle.AESSIV
}

func NewKey(key string) *Key {
	kh, err := subtle.NewAESSIV([]byte(key))
	if err != nil {
		panic(err)
	}

	return &Key{
		kh: kh,
	}
}

func (k *Key) Encrypt(data, additional []byte) ([]byte, error) {
	return k.kh.EncryptDeterministically(data, additional)

}

func (k *Key) Decrypt(ciphertext, additional []byte) ([]byte, error) {
	return k.kh.DecryptDeterministically(ciphertext, additional)
}
