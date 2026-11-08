// Gets all necessary info to put in the display table. Rows configurable through the CLI.
package table

import (
	"github.com/shirou/gopsutil/v3/host"
	"github.com/vlmlee/vanitytable/cmd/format"
)

// Handle OSX, Windows, and Linux

func GetUser() (hostname string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	hostname = info.Hostname
	return
}

func GetOS() {

}

func GetCPU() {

}

func GetGPU() {

}

func GetRAM() {

}

func GetUptime() (uptime string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	uptime = format.FormatUptime(info.Uptime)

	return
}

func GetShell() {

}

func GetKernel() {

}

func GetDistro() {

}

func GetGUITheme() {

}

func GetBattery() {
	
}

func GetTerminalFont() {

}

func GetDisk() {

}