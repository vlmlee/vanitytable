// Gets all necessary info to put in the display table. Rows configurable through the CLI.
package sysinfo

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/vlmlee/vanitytable/cmd/format"
)

// Handle OSX, Windows, and Linux

func GetHostname() (hostname string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	hostname = info.Hostname
	return
}

func GetUser() (username string, err error) {
	currentUser, err := user.Current()

	username = currentUser.Username

	if err != nil {
		return
	}

	return
}

func GetOS() (os string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	switch info.OS {
	case "linux":
		distro, _ := GetDistro()
		os = fmt.Sprintf("%s (%s)", "Linux", distro)
	case "windows":
		os = "Windows"
	case "darwin":
		os = "macOS (darwin)"
	}

	os = os + " " + info.PlatformVersion
	return
}

func GetCPU() (cpuInfo string, err error) {
	info, err := cpu.Info()

	if err != nil {
		return
	}

	cpuInfo = info[0].ModelName

	return
}

func GetGPU() {

}

func GetRAM() (uint64, error) {
	memory, err := mem.VirtualMemory()

	if err != nil {
		return 0, err
	}

	bytesInGB := 1073741824

	return memory.Total / uint64(bytesInGB), nil
}

func GetMemoryUsage() (uint64, error) {
	memory, err := mem.VirtualMemory()

	if err != nil {
		return 0, err
	}

	return uint64(memory.UsedPercent), nil
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

func GetKernel() (kernel string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	kernel = info.KernelArch + " " + info.KernelVersion
	return
}

func GetDistro() (distro string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	distro = info.Platform
	if len(distro) > 0 {
		distro = strings.ToUpper(distro[:1]) + distro[1:]
	}

	distro = fmt.Sprintf("%s %s", distro, info.PlatformVersion)

	return
}

func GetGUITheme() {

}

func GetBattery() {

}

func GetTerminalFont() {

}

func GetDisk() {

}
