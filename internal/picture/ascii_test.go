package picture

import "testing"

func TestGenerateASCIIDefault(t *testing.T) {

}

func TestGenerateASCIIMacOS(t *testing.T) {
	_, err := GenerateASCII("macOS")
	if err != nil {
		t.Error("Could not display macOS ASCII art")
	}
}

func TestGenerateASCIIWindows(t *testing.T) {
	_, err := GenerateASCII("windows")
	if err != nil {
		t.Error("Could not display windows ASCII art")
	}
}

func TestGenerateASCIILinux(t *testing.T) {
	_, err := GenerateASCII("linux")
	if err != nil {
		t.Error("Could not display linux ASCII art")
	}
}
