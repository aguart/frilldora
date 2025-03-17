package gowhisper

type (
	EncryptFunc    func(plaintext, pass []byte) ([]byte, error)
	DecryptFunc    func(ciphertext, pass []byte) ([]byte, error)
	CompressFunc   func(src []byte) ([]byte, error)
	DecompressFunc func(src []byte) ([]byte, error)
)
type Option func(in []byte) ([]byte, error)

func WithEncrypt(pass []byte, enc EncryptFunc) Option {
	return func(in []byte) ([]byte, error) {
		var err error
		in, err = enc(in, pass)
		if err != nil {
			return nil, err
		}
		return in, nil
	}
}

func WithDecrypt(pass []byte, dec DecryptFunc) Option {
	return func(in []byte) ([]byte, error) {
		var err error
		in, err = dec(in, pass)
		if err != nil {
			return nil, err
		}
		return in, nil
	}
}

func WithCompress(comp CompressFunc) Option {
	return func(in []byte) ([]byte, error) {
		var err error
		in, err = comp(in)
		if err != nil {
			return nil, err
		}
		return in, nil
	}
}

func WithDecompress(dec DecompressFunc) Option {
	return func(in []byte) ([]byte, error) {
		var err error
		in, err = dec(in)
		if err != nil {
			return nil, err
		}
		return in, nil
	}
}
