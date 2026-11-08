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

func GetOS() (os string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	os = info.OS
	return
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

func GetKernel() (kernel string, err error)  {
	info, err := host.Info()

	if err != nil {
		return
	}

	kernel = info.KernelArch + " " + info.KernelVersion
	return
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