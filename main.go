package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/alextanhongpin/go-obfuscate/aesgcm"
	"github.com/alextanhongpin/go-obfuscate/aessiv"
	"github.com/alextanhongpin/go-obfuscate/hashid"
)

func main() {
	{
		fmt.Println("aessiv")
		key := aessiv.NewKey(strings.Repeat("A", 64))
		enc, err := key.Encrypt([]byte("hello world"), nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("enc:", base64.URLEncoding.EncodeToString(enc))
		dec, err := key.Decrypt(enc, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("dec:", string(dec))

		fmt.Println("")
	}

	{
		fmt.Println("aesgcm")
		key := aesgcm.NewKey(strings.Repeat("A", 64))
		enc, err := key.Encrypt([]byte("hello world"), nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("enc:", base64.URLEncoding.EncodeToString(enc))
		dec, err := key.Decrypt(enc, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("dec:", string(dec))

		fmt.Println("")
	}

	{
		fmt.Println("hashid")
		key := hashid.NewKey(strings.Repeat("A", 32), 6)
		enc, err := key.Encode(1, 1)
		if err != nil {
			panic(err)
		}

		fmt.Println("enc:", enc)
		dec, err := key.Decode(1, enc)
		if err != nil {
			panic(err)
		}
		fmt.Println("dec:", dec)

		fmt.Println("")
	}
}
