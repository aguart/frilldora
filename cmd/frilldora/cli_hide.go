package main

import (
	"os"

	"github.com/aguart/frilldora"
	"github.com/aguart/frilldora/compress/lzw"
	"github.com/aguart/frilldora/crypto/chacha20"
	"github.com/pkg/errors"
)

func cliHide(opts *cliOptions) error {
	visible, err := os.ReadFile(opts.pathToVisible)
	if err != nil {
		return errors.Wrapf(err, "failed to read file:%s", opts.pathToVisible)
	}
	secret, err := os.ReadFile(opts.pathToSecret)
	if err != nil {
		return errors.Wrapf(err, "failed to read file:%s", opts.pathToVisible)
	}
	var frillOpts []frilldora.Option
	if opts.useCompress {
		frillOpts = append(frillOpts, frilldora.WithCompress(lzw.Compress))
	}
	if len(opts.cryptoPass) > 0 {
		frillOpts = append(frillOpts, frilldora.WithEncrypt([]byte(opts.cryptoPass), chacha20.Encrypt))
	}
	result, err := frilldora.Hide(visible, secret, frillOpts...)
	if err != nil {
		return errors.Wrap(err, "failed to hide")
	}
	if err = os.WriteFile(opts.outputPath, result, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write file:%s", opts.outputPath)
	}
	return nil
}
