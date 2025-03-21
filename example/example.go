package main

import (
	"fmt"

	"github.com/aguart/frilldora"
	"github.com/aguart/frilldora/compress/lzw"
	"github.com/aguart/frilldora/crypto/chacha20"
)

func main() {
	visible := []byte("visible text")
	invisible := []byte("invisible text")

	withHidden, err := frilldora.Hide(visible, invisible)
	if err != nil {
		panic(err)
	}
	secretText, err := frilldora.Reveal(withHidden)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(secretText))
	//output:
	// invisible text

	// Use additional options: crypto and compress
	//
	secretKey := []byte("secret key")
	withHidden, err = frilldora.Hide(visible, invisible,
		frilldora.WithEncrypt(secretKey, chacha20.Encrypt),
		frilldora.WithCompress(lzw.Compress),
	)
	if err != nil {
		panic(err)
	}
	secretText, err = frilldora.Reveal(withHidden,
		frilldora.WithDecompress(lzw.Decompress),
		frilldora.WithDecrypt(secretKey, chacha20.Decrypt),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(secretText))
	//output:
	// invisible text
}
