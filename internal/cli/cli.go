// This package should be used to handle user inputted options and settings
package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DisplayImagePath string
}

var DefaultConfig = &Config{
	DisplayImagePath: "",
}

type Options struct {
	About     bool
	ImagePath string
}

var ProjectDescription = `
This project is a neofetch clone built by vlmlee. 
`

func ParseFlags(args []string, stderr io.Writer) (*Options, error) {
	fs := flag.NewFlagSet("cli", flag.ContinueOnError)
	fs.SetOutput(stderr)

	opts := &Options{}

	fs.StringVar(&opts.ImagePath, "imagePath", "", "Set the display image path")
	fs.BoolVar(&opts.About, "about", false, "Info about vanitytable")

	if err := fs.Parse(args); err != nil {
		return opts, err
	}

	return opts, nil
}

// Prints info about this project to stdout
func PrintAbout() {
	fmt.Println(ProjectDescription)
}

func SetDisplayPictureFilePath() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	f := flag.Lookup("imagePath")

	if f == nil {
		return fmt.Errorf("No imagePath supplied")
	}

	config.DisplayImagePath = f.Value.String()
	SaveConfig(config)

	return nil
}

func configFilePath() (string, error) {
	baseDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, "vanity_table_config.json"), nil
}

func SaveConfig(config *Config) error {
	path, err := configFilePath()

	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o600)
}

func LoadConfig() (*Config, error) {
	path, err := configFilePath()

	if err != nil {
		return DefaultConfig, err
	}

	data, err := os.ReadFile(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return DefaultConfig, nil
		}
	}

	config := DefaultConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return DefaultConfig, nil
	}

	return config, nil
}
