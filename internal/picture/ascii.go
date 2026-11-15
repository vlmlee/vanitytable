package picture

import (
	"embed"
	_ "embed"
)

//go:embed *.txt
var textFiles embed.FS

func GenerateASCII() {

}

func ASCIIMacOS() (string, error) {
	content, err := textFiles.ReadFile("macOS.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ASCIIWindows() (string, error) {
	content, err := textFiles.ReadFile("windows.txt")
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ASCIILinux() {

}
