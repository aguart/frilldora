# Frilldora
#### Enables to use cloaking.  
<p align="center">
  <img src="https://github.com/aguart/frilldora/blob/main/4088.png" title="frilldora card">
</p> 
  
### Prologue
In a very old game, Ragnarok Online, there was a small chance to loot a Frilldora Card. This card allowed the player to become invisible.  
This service does the same thing. But don’t take it too seriously—it's more of a joke library. It allows you to use Unicode "gray areas" to hide invisible text behind visible characters.

### Description
- You can hide secret text behind visible text
- You can reveal secret text
- You can encrypt your secret text (and decrypt)
- You can compress your secret text (and decompress)
- You can encrypt and compress together

_Be careful. For right work, if you use order: crypt -> compress, when reveal you should use: decompress -> decrypt. It's your responsibility_

### For what
- Play the spy game
- Make watermarks for texts
- NOT for serious stuff

### Install CLI tool
```
go install github.com/aguart/frilldora/cmd/frilldora@latest
```
### Commands
```
USAGE:
   Frilldora [global options] [command [command options]]

COMMANDS:
   hide     hide secret text
   reveal   reveal secret text
   help, h  Shows a list of commands or help for one command
-----
USAGE:
   Frilldora hide [command [command options]]

OPTIONS:
   --visible value, -v value       path to file with visible text
   --secret value, -s value        path to file with secret text
   --output value, -o value        path to output file
   --compress, -c                  is compress secret text (default: false)
   --crypto-key value, --ck value  key used to encrypt secret text
   --help, -h                      show help
------
USAGE:
   Frilldora reveal [command [command options]]

OPTIONS:
   --input value, -i value         path to input file
   --output value, -o value        path to output file
   --decompress, --dc              is decompress secret text (default: false)
   --crypto-key value, --ck value  key used to decrypt secret text
   --help, -h                      show help

```
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