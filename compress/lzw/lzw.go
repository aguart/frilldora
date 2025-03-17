package lzw

import (
	"bytes"
	"compress/lzw"
	"io"
)

const litw = 8

func Compress(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	wr := lzw.NewWriter(&buf, lzw.LSB, litw)
	if _, err := wr.Write(src); err != nil {
		return nil, err
	}
	_ = wr.Close()
	return buf.Bytes(), nil
}

func Decompress(src []byte) ([]byte, error) {
	rr := lzw.NewReader(bytes.NewReader(src), lzw.LSB, litw)
	var decompressedData bytes.Buffer
	if _, err := io.Copy(&decompressedData, rr); err != nil {
		return nil, err
	}
	_ = rr.Close()
	return decompressedData.Bytes(), nil
}
