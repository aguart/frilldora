package lzw

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressDecompress(t *testing.T) {
	cases := []struct {
		descr string
		data  []byte
	}{
		{
			"ok",
			[]byte(`qwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,
qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp,
qqwweerrttyyuuiiooppp,qqwweerrttyyuuiiooppp`),
		},
		{
			"empty data",
			[]byte(""),
		},
		{
			"with long characters",
			[]byte("vi世si界ble"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.descr, func(t *testing.T) {
			comp, err := Compress(tc.data)
			require.NoError(t, err)
			decomp, err := Decompress(comp)
			require.NoError(t, err)
			require.Equal(t, tc.data, decomp)
		})
	}
}
