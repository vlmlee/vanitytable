package main

import (
	"errors"
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
	opts, err := cli.ParseFlags(args, os.Stderr)
	if err != nil {
		return err
	}

	if opts.About {
		cli.PrintAbout()
	}

	if opts.ImagePath != "" {
		if err := cli.SetDisplayPictureFilePath(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}
	}

	return render.RenderTable()
}
