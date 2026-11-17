package picture

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed *.txt
var textFiles embed.FS

func GenerateASCII(os string) (string, error) {
	filename := ""
	switch os {
	case "macOS":
		filename = "macOS.txt"
	case "linux":
		filename = "linux.txt"
	case "windows":
		filename = "windows.txt"
	default:
		return "", fmt.Errorf("No os provided")
	}

	content, err := textFiles.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
