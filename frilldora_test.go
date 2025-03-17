package gowhisper

import (
	"testing"

	"github.com/aguart/frilldora/compress/lzw"
	"github.com/aguart/frilldora/crypto/chacha20"
	"github.com/stretchr/testify/require"
)

func TestWork(t *testing.T) {
	cases := []struct {
		desc      string
		visible   []byte
		invisible []byte
		encOpt    Option
		decOpt    Option
		compOpt   Option
		decompOpt Option
	}{
		{
			desc:      "ok",
			visible:   []byte("visible"),
			invisible: []byte("invisible"),
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		{
			desc:      "ok with crypt",
			visible:   []byte("visible"),
			invisible: []byte("invisible"),
			encOpt:    WithEncrypt([]byte("pass"), chacha20.Encrypt),
			decOpt:    WithDecrypt([]byte("pass"), chacha20.Decrypt),
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		{
			desc:      "ok with compress",
			visible:   []byte("visible"),
			invisible: []byte("invisible"),
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   WithCompress(lzw.Compress),
			decompOpt: WithDecompress(lzw.Decompress),
		},
		{
			desc:      "ok with crypt and compress",
			visible:   []byte("visible"),
			invisible: []byte("invisible"),
			encOpt:    WithEncrypt([]byte("pass"), chacha20.Encrypt),
			decOpt:    WithDecrypt([]byte("pass"), chacha20.Decrypt),
			compOpt:   WithCompress(lzw.Compress),
			decompOpt: WithDecompress(lzw.Decompress),
		},
	}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			withHidden, err := Hide(tc.visible, tc.invisible, tc.compOpt, tc.encOpt)
			require.NoError(t, err)
			t.Log(len(withHidden))
			hiddenText, err := Reveal(withHidden, tc.decOpt, tc.decompOpt)
			require.NoError(t, err)
			t.Log(len(hiddenText))
			require.Equal(t, tc.invisible, hiddenText)
		})
	}
}
