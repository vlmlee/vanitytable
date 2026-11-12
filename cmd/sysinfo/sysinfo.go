// Gets all necessary info to put in the display table. Rows configurable through the CLI.
package sysinfo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/vlmlee/vanitytable/cmd/format"
)

const bytesInGB = 1073741824

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

func GetGPU() string {
	switch runtime.GOOS {
	case "darwin":
		out, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
		if err == nil {
			for _, line := range strings.Split(string(out), "\n") {
				if strings.Contains(line, "Chipset Model:") {
					parts := strings.SplitN(line, ":", 2)
					return strings.TrimSpace(parts[1])
				}
			}
		}

	case "linux":
		out, err := exec.Command("sh", "-c", "lspci | grep -iE 'vga|3d|display'").Output()
		if err == nil {
			lines := strings.Split(strings.TrimSpace(string(out)), "\n")
			if len(lines) > 0 && lines[0] != "" {
				parts := strings.SplitN(lines[0], ": ", 2)
				if len(parts) == 2 {
					return strings.TrimSpace(parts[1])
				}
			}
		}

	case "windows":
		cmd := exec.Command("powershell", "-NoProfile", "-Command", "(Get-CimInstance Win32_VideoController).Name")
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err == nil {
			lines := strings.Split(strings.TrimSpace(out.String()), "\n")
			if len(lines) > 0 && lines[0] != "" {
				return strings.TrimSpace(lines[0])
			}
		}
	}

	return "unknown"
}

func GetRAM() (uint64, error) {
	memory, err := mem.VirtualMemory()

	if err != nil {
		return 0, err
	}

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

func GetShell() string {
	if runtime.GOOS == "windows" {
		if os.Getenv("PSModulePath") != "" {
			return "PowerShell"
		}
		if comspec := os.Getenv("COMSPEC"); comspec != "" {
			return filepath.Base(comspec)
		}
		return "unknown"
	}

	if shell := os.Getenv("SHELL"); shell != "" {
		return filepath.Base(shell)
	}
	return "unknown"
}

func GetKernel() (kernel string, err error) {
	info, err := host.Info()

	if err != nil {
		return
	}

	kernel = fmt.Sprintf("%s (%s)", info.KernelVersion, info.KernelArch)
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

func GetTerminal() string {
	if termProgram := os.Getenv("TERM_PROGRAM"); termProgram != "" {
		if termProgram == "Apple_Terminal" {
			return "Terminal"
		}
		if termProgram == "iTerm.app" {
			return "iTerm2"
		}
		return termProgram
	}

	if os.Getenv("WT_SESSION") != "" {
		return "Windows Terminal"
	}

	knownShells := map[string]bool{
		"zsh": true, "bash": true, "fish": true, "sh": true,
		"pwsh": true, "powershell.exe": true, "cmd.exe": true, "go": true,
	}

	pid := int32(os.Getppid())
	for pid > 1 {
		proc, err := process.NewProcess(pid)
		if err != nil {
			break
		}

		name, err := proc.Name()
		if err != nil {
			break
		}

		lowerName := strings.ToLower(name)
		if !knownShells[lowerName] {
			return name
		}

		ppid, err := proc.Ppid()
		if err != nil || ppid == pid {
			break
		}
		pid = ppid
	}

	if term := os.Getenv("TERM"); term != "" {
		return term
	}

	return "unknown"
}

func GetDisk() (diskspace string, err error) {
	// Use C:\ on Windows, / on macOS and Linux
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:\\"
	}

	usage, err := disk.Usage(path)
	if err != nil {
		return
	}

	diskspace = fmt.Sprintf("%.2f GB ~ %.1f%% used", float64(usage.Total)/bytesInGB, usage.UsedPercent)

	return
}
