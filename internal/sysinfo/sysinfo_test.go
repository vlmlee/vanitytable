package sysinfo

import (
	"strings"
	"testing"
)

func TestTable(t *testing.T) {

}

func TestGetHostname(t *testing.T) {
	if _, err := GetHostname(); err != nil {
		t.Error("Hostname doesn't exist")
	}
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
	cpu, err := GetCPU()

	if err != nil || cpu == "" {
		t.Error("Could not get CPU")
	}
}

func TestGetGPU(t *testing.T) {
	gpu := GetGPU()

	if gpu == "unknown" {
		t.Error("Could not get GPU")
	}
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
	shell := GetShell()

	if shell == "unknown" {
		t.Error("Could not get shell")
	}
}

func TestGetKernel(t *testing.T) {
	kernel, err := GetKernel()
	if err != nil || kernel == " " {
		t.Error("Could not get the kernel")
	}
}

func TestGetDistro(t *testing.T) {
	distro, err := GetDistro()

	if err != nil || distro == "" {
		t.Error("Could not get distro")
	}
}

func TestGetGUITheme(t *testing.T) {

}

func TestGetTerminalFont(t *testing.T) {
	terminal := GetTerminal()

	if terminal == "unknown" {
		t.Error("Could not get terminal")
	}
}

func TestGetDisk(t *testing.T) {
	_, err := GetDisk()

	if err != nil {
		t.Error("Could not get disk space")
	}
}
