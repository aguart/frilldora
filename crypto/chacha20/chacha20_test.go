package chacha20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	cases := []struct {
		description string
		pass        []byte
		data        []byte
		expected    []byte
	}{
		{
			"ok",
			[]byte("pass"),
			[]byte("data"),
			[]byte("data"),
		},
		{
			"empty pass",
			[]byte(""),
			[]byte("data2"),
			[]byte("data2"),
		},
		{
			"empty data",
			[]byte("pass"),
			[]byte(""),
			[]byte(""),
		},
		{
			"nil data",
			[]byte("pass"),
			nil,
			[]byte(""),
		},
		{
			"nil pass",
			nil,
			[]byte("data"),
			[]byte("data"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			encrypted, err := Encrypt(tc.data, tc.pass)
			require.NoError(t, err)
			decrypted, err := Decrypt(encrypted, tc.pass)
			require.NoError(t, err)
			require.Equal(t, tc.expected, decrypted)
		})
	}

}
