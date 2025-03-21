package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var (
		pathToVisible string
		pathToSecret  string
		outputPath    string
		inputPath     string
	)
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
						Destination: &pathToVisible,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "secret",
						Aliases:     []string{"s"},
						Usage:       "path to file with secret text",
						Destination: &pathToSecret,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "path to output file",
						Destination: &outputPath,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(pathToVisible, pathToSecret, outputPath)
					return nil
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
						Destination: &inputPath,
					},
					&cli.StringFlag{
						Required:    true,
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "path to output file",
						Destination: &outputPath,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(inputPath, outputPath)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
