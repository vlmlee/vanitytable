package picture

import "testing"

func TestGenerateASCIIDefault(t *testing.T) {

}

func TestGenerateASCIIMacOS(t *testing.T) {
	_, err := ASCIIMacOS()
	if err != nil {
		t.Error("Could not display macOS ASCII art")
	}
}

func TestGenerateASCIIWindows(t *testing.T) {
	_, err := ASCIIWindows()
	if err != nil {
		t.Error("Could not display windows ASCII art")
	}
}

func TestGenerateASCIILinux(t *testing.T) {

}
