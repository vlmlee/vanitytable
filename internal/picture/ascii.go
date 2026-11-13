package picture

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getArt(name string) (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine source file path")
	}

	path := filepath.Join(filepath.Dir(currentFile), name)

	content, err := os.ReadFile(path)
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
