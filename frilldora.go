package gowhisper

import (
	"bytes"
	"io"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	// Variation selectors block https://unicode.org/charts/nameslist/n_FE00.html
	// VS1--VS16
	variationSelectorStart = 0xFE00
	variationSelectorEnd   = 0xFE0F

	// Variation selectors supplement https://unicode.org/charts/nameslist/n_E0100.html
	// VS17--VS256
	variationSelectorSupplementStart = 0xE0100
	variationSelectorSupplementEnd   = 0xE01EF
)

func Hide(visible, invisible []byte, opts ...Option) ([]byte, error) {
	for _, opt := range opts {
		var err error
		invisible, err = opt(invisible)
		if err != nil {
			return nil, err
		}
	}
	invReader := bytes.NewReader(invisible)
	result := new(bytes.Buffer)
	result.Grow(len(visible) + len(invisible))
	for _, charVis := range string(visible) {
		result.WriteRune(charVis)
		bb, err := invReader.ReadByte()
		if errors.Is(err, io.EOF) {
			continue
		}
		if err != nil {
			return nil, err
		}
		result.WriteRune(toVariationSelector(bb))
	}
	for {
		bb, err := invReader.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		}
		result.WriteRune(toVariationSelector(bb))
	}

	return result.Bytes(), nil
}

func Reveal(input []byte, opts ...Option) ([]byte, error) {
	var decodedBytes []byte
	for i, w := 0, 0; i < len(input); i += w {
		r, width := utf8.DecodeRune(input[i:])
		w = width

		if byteValue, ok := fromVariationSelector(r); ok {
			decodedBytes = append(decodedBytes, byteValue)
		}
	}

	for _, opt := range opts {
		var err error
		decodedBytes, err = opt(decodedBytes)
		if err != nil {
			return nil, err
		}
	}

	return decodedBytes, nil
}

// toVariationSelector make hidden variant selector rune from byte
func toVariationSelector(b byte) rune {
	if b < 16 {
		return rune(variationSelectorStart + int(b))
	}
	return rune(variationSelectorSupplementStart + int(b) - 16)
}

// fromVariationSelector get byte from hidden variant selector
func fromVariationSelector(codePoint rune) (byte, bool) {
	if codePoint >= variationSelectorStart && codePoint <= variationSelectorEnd {
		return byte(codePoint - variationSelectorStart), true
	} else if codePoint >= variationSelectorSupplementStart && codePoint <= variationSelectorSupplementEnd {
		return byte(codePoint - variationSelectorSupplementStart + 16), true
	}
	return 0, false
}
