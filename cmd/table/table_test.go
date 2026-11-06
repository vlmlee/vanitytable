package table

import (
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	
}

func TestGetUser(t *testing.T) {

}

func TestGetOS(t *testing.T) {

}

func TestGetCPU(t *testing.T) {

}

func TestGetGPU(t *testing.T) {

}

func TestGetRAM(t *testing.T) {

}

func TestGetUptime(t *testing.T) {
	if !strings.ContainsAny(GetUptime(), "123456789") {
		t.Error("Could not get the uptime")
	}
}

func TestGetShell(t *testing.T) {

}

func TestGetKernel(t *testing.T) {

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