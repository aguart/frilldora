package main

import (
	"os"

	"github.com/aguart/frilldora"
	"github.com/aguart/frilldora/compress/lzw"
	"github.com/aguart/frilldora/crypto/chacha20"
	"github.com/pkg/errors"
)

func cliReveal(opts *cliOptions) error {
	input, err := os.ReadFile(opts.inputPath)
	if err != nil {
		return errors.Wrapf(err, "failed to read file:%s", opts.inputPath)
	}
	var frillOpts []frilldora.Option
	if len(opts.cryptoPass) > 0 {
		frillOpts = append(frillOpts, frilldora.WithDecrypt([]byte(opts.cryptoPass), chacha20.Decrypt))
	}
	if opts.useCompress {
		frillOpts = append(frillOpts, frilldora.WithDecompress(lzw.Decompress))
	}
	result, err := frilldora.Reveal(input, frillOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to reveal")
	}
	if err = os.WriteFile(opts.outputPath, result, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write file:%s", opts.outputPath)
	}
	return nil
}
