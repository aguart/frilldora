package frilldora

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzWork(f *testing.F) {
	testcases := []struct {
		visible   []byte
		invisible []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("a"), []byte("b")},
		{[]byte("visible"), []byte("invisible")},
		{[]byte("visible\x00"), []byte("invisible")},
	}
	for _, tc := range testcases {
		f.Add(tc.visible, tc.invisible)
	}
	f.Fuzz(func(t *testing.T, visible []byte, invisible []byte) {
		t.Logf("%X\n", visible)
		t.Logf("%X\n", invisible)
		withHidden, err := Hide(visible, invisible)
		require.NoError(t, err)
		hidden, err := Reveal(withHidden)
		require.NoError(t, err)
		require.Equal(t, invisible, hidden)
	})
}
