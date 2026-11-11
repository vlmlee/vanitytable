package table

import (
	"strings"
	"testing"
)

func TestTable(t *testing.T) {

}

func TestGetUser(t *testing.T) {
	if _, err := GetUser(); err != nil {
		t.Error("User doesn't exist")
	}
}

func TestGetOS(t *testing.T) {
	os, err := GetOS()
	if err != nil || os == "" {
		t.Error("Could not get the OS")
	}
}

func TestGetCPU(t *testing.T) {

}

func TestGetGPU(t *testing.T) {

}

func TestGetRAM(t *testing.T) {
	totalRAM, err := GetRAM()

	if totalRAM == 0 || err != nil {
		t.Error("could not get RAM")
	}
}

func TestGetMemoryUsage(t *testing.T) {
	usage, err := GetMemoryUsage()

	if usage == 0 || err != nil {
		t.Error("could not get the memory usage")
	}
}

func TestGetUptime(t *testing.T) {
	uptime, err := GetUptime()
	if !strings.ContainsAny(uptime, "123456789") || err != nil {
		t.Error("Could not get the uptime")
	}
}

func TestGetShell(t *testing.T) {

}

func TestGetKernel(t *testing.T) {
	kernel, err := GetKernel()
	if err != nil || kernel == " " {
		t.Error("Could not get the kernel")
	}
}

func TestGetDistro(t *testing.T) {

}

func TestGetGUITheme(t *testing.T) {

}

func TestGetBattery(t *testing.T) {

}

func TestGetTerminalFont(t *testing.T) {

}

func TestGetDisk(t *testing.T) {

}
