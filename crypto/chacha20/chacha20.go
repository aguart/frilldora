package chacha20

import (
	cryptorand "crypto/rand"
	"crypto/sha256"

	"github.com/pkg/errors"
	"golang.org/x/crypto/chacha20poly1305"
)

func Encrypt(plaintext, pass []byte) ([]byte, error) {
	key := sha256.Sum256(pass)
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(plaintext)+aead.Overhead())
	if _, err = cryptorand.Read(nonce); err != nil {
		return nil, err
	}
	encryptedMsg := aead.Seal(nonce, nonce, plaintext, nil)
	return encryptedMsg, nil
}

func Decrypt(encryptedMsg, pass []byte) ([]byte, error) {
	key := sha256.Sum256(pass)
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, err
	}
	if len(encryptedMsg) < aead.NonceSize() {
		return nil, errors.New("invalid encrypted message")
	}
	nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	if plaintext == nil {
		return []byte(""), nil
	}

	return plaintext, nil
}
