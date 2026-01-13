package main

import (
	"context"
	"fmt"
	"os"

	ps "code"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Usage:   "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return fmt.Errorf("path argument is required")
			}
			path := cmd.Args().First()
			result, err := ps.GetPathSize(path, false, cmd.Bool("human"), false)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
