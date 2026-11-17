package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vlmlee/vanitytable/internal/cli"
	"github.com/vlmlee/vanitytable/internal/render"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		log.Fatalf("Error: %v\n", err)
	}
}

func run(args []string) error {
	fs, opts, err := cli.ParseFlags(args, os.Stderr)
	if err != nil {
		return err
	}

	if opts.About {
		cli.PrintAbout()
		os.Exit(0)
	}

	if opts.Image != "" {
		if err := cli.SetDisplayPictureFilePath(fs); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}
	}

	return render.RenderTable()
}
