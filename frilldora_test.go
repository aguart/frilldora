package frilldora

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
			desc:      "ok empty invisible",
			visible:   []byte("visible 02"),
			invisible: []byte(""),
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		//{
		//	desc:      "ok empty visible",
		//	visible:   []byte(""),
		//	invisible: []byte("0"),
		//	encOpt:    func(in []byte) ([]byte, error) { return in, nil },
		//	decOpt:    func(in []byte) ([]byte, error) { return in, nil },
		//	compOpt:   func(in []byte) ([]byte, error) { return in, nil },
		//	decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		//},
		//{
		//	desc:      "ok from fuzz 01",
		//	visible:   []byte{0xEF, 0xB8, 0x86},
		//	invisible: []byte{0x30},
		//	encOpt:    func(in []byte) ([]byte, error) { return in, nil },
		//	decOpt:    func(in []byte) ([]byte, error) { return in, nil },
		//	compOpt:   func(in []byte) ([]byte, error) { return in, nil },
		//	decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		//},
		{
			desc:      "ok from fuzz 02",
			visible:   []byte{0xEF, 0xB8, 0x8F},
			invisible: []byte{},
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		{
			desc:      "ok invisible much longer",
			visible:   []byte("visible"),
			invisible: []byte("invisible.invisible.invisible.invisible.invisible.invisible"),
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		{
			desc:      "ok long unicode visible chars",
			visible:   []byte("vi世si界ble"),
			invisible: []byte("invisible.invisible.invisible.invisible.invisible.invisible"),
			encOpt:    func(in []byte) ([]byte, error) { return in, nil },
			decOpt:    func(in []byte) ([]byte, error) { return in, nil },
			compOpt:   func(in []byte) ([]byte, error) { return in, nil },
			decompOpt: func(in []byte) ([]byte, error) { return in, nil },
		},
		{
			desc:      "ok long unicode invisible chars",
			visible:   []byte("vi世si界ble"),
			invisible: []byte("invi世si界ble.invi世si界ble"),
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
		{
			desc:      "ok with long unicode chars, with crypt and compress",
			visible:   []byte("vi世si界ble"),
			invisible: []byte("invi世si界ble.invi世si界ble"),
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
			//t.Log(string(withHidden))
			//t.Log(len(withHidden))
			hiddenText, err := Reveal(withHidden, tc.decOpt, tc.decompOpt)
			require.NoError(t, err)
			//t.Log(len(hiddenText))
			require.Equal(t, tc.invisible, hiddenText)
		})
	}
}

func BenchmarkHide(b *testing.B) {
	visible := []byte("visible")
	invisible := []byte("invisible.invisible.invisible.invisible.invisible.invisible.invisible.invisible")
	b.Run("hide", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = Hide(visible, invisible)
		}
	})
}
