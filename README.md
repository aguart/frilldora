# Frilldora
#### Enables to use cloaking.  
<p align="center">
  <img src="https://github.com/aguart/frilldora/blob/main/4088.png" title="frilldora card">
</p> 
  
### Prologue
In a very old game, Ragnarok Online, there was a small chance to loot a Frilldora Card. This card allowed the player to become invisible.  
This service does the same thing. But don’t take it too seriously—it's more of a joke library. It allows you to use Unicode "gray areas" to hide invisible text behind visible characters.

### Description
description

### Example 
```go
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
```