package format

import (
	"fmt"
	"strings"
)

func FormatTable() {

}

func FormatDisplayPicture() {

}

func FormatString() {
	
}

func FormatUptime(uptime uint64) string {
	days := uptime / 86400
	hours := (uptime % 86400) / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60

	var parts []string

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 || days > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 || hours > 0 || days > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	parts = append(parts, fmt.Sprintf("%ds", seconds)) 

	return strings.Join(parts, " ")
}