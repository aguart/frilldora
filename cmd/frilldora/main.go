package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

type cliOptions struct {
	pathToVisible string
	pathToSecret  string
	outputPath    string
	inputPath     string
	useCompress   bool
	cryptoPass    string
}

func main() {
	opts := cliOptions{}

	cmd := &cli.Command{
		Name:  "Frilldora",
		Usage: "cloaking tool",
		Commands: []*cli.Command{
			{
				Name:  "hide",
				Usage: "hide secret text",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Required:    true,
						Name:        "visible",
						Aliases:     []string{"v"},
						Usage:       "path to file with visible text",
						Destination: &opts.pathToVisible,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "secret",
						Aliases:     []string{"s"},
						Usage:       "path to file with secret text",
						Destination: &opts.pathToSecret,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "path to output file",
						Destination: &opts.outputPath,
					},
					&cli.BoolFlag{
						Name:        "compress",
						Aliases:     []string{"c"},
						Usage:       "is compress secret text",
						Destination: &opts.useCompress,
					},
					&cli.StringFlag{
						Name:        "crypto-key",
						Aliases:     []string{"ck"},
						Usage:       "key used to encrypt secret text",
						Destination: &opts.cryptoPass,
					},
				},
				Action: func(ctx context.Context, _ *cli.Command) error {
					return cliHide(&opts)
				},
			},
			{
				Name:  "reveal",
				Usage: "reveal secret text",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Required:    true,
						Name:        "input",
						Aliases:     []string{"i"},
						Usage:       "path to input file",
						Destination: &opts.inputPath,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "path to output file",
						Destination: &opts.outputPath,
					},
					&cli.BoolFlag{
						Name:        "decompress",
						Aliases:     []string{"dc"},
						Usage:       "is decompress secret text",
						Destination: &opts.useCompress,
					},
					&cli.StringFlag{
						Name:        "crypto-key",
						Aliases:     []string{"ck"},
						Usage:       "key used to decrypt secret text",
						Destination: &opts.cryptoPass,
					},
				},
				Action: func(ctx context.Context, _ *cli.Command) error {
					return cliReveal(&opts)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
