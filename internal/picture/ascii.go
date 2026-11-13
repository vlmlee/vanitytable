package picture

import (
	_ "embed"
	"os"
	"path/filepath"
)

func getArt(name string) (string, error) {
	wd, err := os.Getwd()
	filePath := filepath.Join(wd, name)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GenerateASCII() {

}

func ASCIIMacOS() (string, error) {
	ascii, err := getArt("macOS.txt")
	if err != nil {
		return "", err
	}

	return ascii, nil
}

func ASCIIWindows() (string, error) {
	ascii, err := getArt("windows.txt")
	if err != nil {
		return "", err
	}

	return ascii, nil
}

func ASCIILinux() {

}
