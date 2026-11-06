// Gets all necessary info to put in the display table. Rows configurable through the CLI.
package table

import (
	"fmt"
	"time"
)

// Handle OSX, Windows, and Linux

func GetUser() {

}

func GetOS() {

}

func GetCPU() {

}

func GetGPU() {

}

func GetRAM() {

}

func GetUptime() string {
	bootTime := time.Since(time.Now().Add(-1 * time.Hour)) 
	return fmt.Sprintf("%d mins", int(bootTime.Minutes()))
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