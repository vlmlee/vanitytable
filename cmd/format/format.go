package format

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/image/draw"
)

type TableInfo struct {
	Cols, Rows   int
	CellW, CellH float64
}

func GetTerminalInfo() TableInfo {
	ws := &struct {
		Row, Col, Xpixel, YPixel uint16
	}{}

	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	info := TableInfo{
		Cols:  80,
		Rows:  24,
		CellW: 10.0,
		CellH: 20.0,
	}

	if err == 0 && ws.Col > 0 && ws.Row > 0 {
		info.Cols = int(ws.Col)
		info.Rows = int(ws.Row)
		if ws.Xpixel > 0 && ws.YPixel > 0 {
			info.CellW = float64(ws.Xpixel) / float64(ws.Col)
			info.CellH = float64(ws.YPixel) / float64(ws.Row)
		}
	}

	return info
}

func FormatDisplayPicture(img image.Image, info TableInfo, maxCols, maxRows int) image.Image {
	srcW := img.Bounds().Dx()
	srcH := img.Bounds().Dy()

	maxPxW := int(float64(maxCols) * info.CellW)
	maxPxH := int(float64(maxRows) * info.CellH)

	if srcW <= maxPxW && srcH <= maxPxH {
		return img
	}

	scaleW := float64(maxPxW) / float64(srcW)
	scaleH := float64(maxPxH) / float64(srcH)
	scale := math.Min(scaleW, scaleH)

	dstW := max(1, int(math.Round(float64(srcW)*scale)))
	dstH := max(1, int(math.Round(float64(srcH)*scale)))

	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

func FormatTable() {

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
