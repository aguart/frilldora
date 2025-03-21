package frilldora

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
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
	visible = []byte(Clean(string(visible)))
	for _, opt := range opts {
		var err error
		invisible, err = opt(invisible)
		if err != nil {
			return nil, err
		}
	}

	result := new(bytes.Buffer)
	runesVisible := []rune(string(visible))
	lenRunes, lenBytes := len(runesVisible), len(invisible)
	result.Grow(len(visible) + len(invisible)*4) // multiple 4 because hidden byte convert to rune(4bytes)

	longerLen, shorterLen := lenRunes, lenBytes
	isRunesLonger := true
	if lenBytes > lenRunes {
		longerLen, shorterLen = lenBytes, lenRunes
		isRunesLonger = false
	}

	step := float64(longerLen) / float64(shorterLen)
	longerIdx, shorterIdx := 0, 0

	for shorterIdx < shorterLen || longerIdx < longerLen {
		nextLongerIdx := int(float64(shorterIdx+1) * step)
		if nextLongerIdx > longerLen {
			nextLongerIdx = longerLen
		}

		if isRunesLonger {
			for ; longerIdx < nextLongerIdx; longerIdx++ {
				result.WriteRune(runesVisible[longerIdx])
			}
		} else {
			for ; longerIdx < nextLongerIdx; longerIdx++ {
				result.WriteRune(toVariationSelector(invisible[longerIdx]))
			}
		}

		if shorterIdx < shorterLen {
			if isRunesLonger {
				result.WriteRune(toVariationSelector(invisible[shorterIdx]))
			} else {
				result.WriteRune(runesVisible[shorterIdx])
			}
			shorterIdx++
		}
	}

	return result.Bytes(), nil
}

func Reveal(input []byte, opts ...Option) ([]byte, error) {
	decodedBytes := make([]byte, 0)
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

func Clean(s string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsGraphic(r) {
			return -1
		}
		if !unicode.IsPrint(r) {
			return -1
		}
		if unicode.Is(InvisibleRanges, r) {
			return -1
		}
		return r
	}, s)
}

var InvisibleRanges = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 11, Hi: 13, Stride: 1},
		{Lo: 127, Hi: 160, Stride: 33},
		{Lo: 173, Hi: 847, Stride: 674},
		{Lo: 1564, Hi: 4447, Stride: 2883},
		{Lo: 4448, Hi: 6068, Stride: 1620},
		{Lo: 6069, Hi: 6155, Stride: 86},
		{Lo: 6156, Hi: 6158, Stride: 1},
		{Lo: 7355, Hi: 7356, Stride: 1},
		{Lo: 8192, Hi: 8207, Stride: 1},
		{Lo: 8234, Hi: 8239, Stride: 1},
		{Lo: 8287, Hi: 8303, Stride: 1},
		{Lo: 10240, Hi: 12288, Stride: 2048},
		{Lo: 12644, Hi: 65024, Stride: 52380},
		{Lo: 65025, Hi: 65039, Stride: 1},
		{Lo: 65279, Hi: 65440, Stride: 161},
		{Lo: 65520, Hi: 65528, Stride: 1},
		{Lo: 65532, Hi: 65532, Stride: 1},
	},
	R32: []unicode.Range32{
		{Lo: 78844, Hi: 119155, Stride: 40311},
		{Lo: 119156, Hi: 119162, Stride: 1},
		{Lo: 917504, Hi: 917631, Stride: 1},
		{Lo: 917760, Hi: 917999, Stride: 1},
	},
	LatinOffset: 2,
}
