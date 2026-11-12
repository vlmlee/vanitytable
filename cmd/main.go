package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/dolmen-go/kittyimg"
	"github.com/vlmlee/vanitytable/cmd/cli"
	"github.com/vlmlee/vanitytable/cmd/format"
	"github.com/vlmlee/vanitytable/cmd/sysinfo"
)

var imagePathFlag string
var printAboutFlag bool

func init() {
	flag.StringVar(&imagePathFlag, "imagePath", "", "Set the display image path")
	flag.BoolVar(&printAboutFlag, "about", false, "Info about vanitytable")
}

func main() {
	flag.Parse()

	if printAboutFlag {
		cli.PrintAbout()
		return
	}

	if imagePathFlag != "" {
		cli.SetDisplayPictureFilePath()
		return
	}

	hostname, _ := sysinfo.GetHostname()
	user, _ := sysinfo.GetUser()
	osName, _ := sysinfo.GetOS()
	kernel, _ := sysinfo.GetKernel()
	cpuName, _ := sysinfo.GetCPU()
	gpuName := sysinfo.GetGPU()
	diskspace, _ := sysinfo.GetDisk()
	ram, _ := sysinfo.GetRAM()
	memUsage, _ := sysinfo.GetMemoryUsage()
	terminal := sysinfo.GetTerminal()
	shell := sysinfo.GetShell()
	uptime, _ := sysinfo.GetUptime()

	f, err := os.Open("/Users/michaellee/Desktop/vanitytable/cmd/akari.jpeg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer f.Close()

	origImg, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(format.Glossy_Pink)).
		Background(lipgloss.Color(format.Fantasy_Blue)).
		Padding(0, 1).Align(lipgloss.Center)

	cellStyle := lipgloss.NewStyle().Padding(0, 1)

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderRow(true).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).
		Headers("SYSTEM", "INFO").
		Rows(
			[]string{"Host", hostname},
			[]string{"User", user},
			[]string{"OS", osName},
			[]string{"Kernel", kernel},
			[]string{"CPU", cpuName},
			[]string{"GPU", gpuName},
			[]string{"Disk", diskspace},
			[]string{"RAM", strconv.FormatUint(ram, 10) + "GB ~ " + strconv.FormatUint(memUsage, 10) + "% used"},
			[]string{"Terminal", terminal},
			[]string{"Shell", shell},
			[]string{"Uptime", uptime},
		)

	renderedTable := lipgloss.NewStyle().PaddingRight(1).PaddingBottom(1).Render(t.Render())
	tableLines := strings.Split(renderedTable, "\n")
	tableWidth := lipgloss.Width(renderedTable)

	info := format.GetTerminalInfo()
	leftPad := 2
	gap := 2
	maxImgCols := max(10, info.Cols-tableWidth-leftPad-gap)
	maxImgRows := max(5, info.Rows-2)
	imgTopPad := 1

	img := format.FormatDisplayPicture(origImg, info, maxImgCols, maxImgRows)

	imgCols := int(math.Ceil(float64(img.Bounds().Dx()) / info.CellW))
	imgRows := int(math.Ceil(float64(img.Bounds().Dy()) / info.CellH))

	totalHeight := max(imgRows+imgTopPad, len(tableLines))
	tableOffsetCol := leftPad + imgCols + gap

	fmt.Print("\n")

	fmt.Print(strings.Repeat("\n", totalHeight))
	fmt.Printf("\033[%dA\r", totalHeight)

	if imgTopPad > 0 {
		fmt.Printf("\033[%dB", imgTopPad)
	}

	if leftPad > 0 {
		fmt.Printf("\033[%dC", leftPad)
	}

	if err := kittyimg.Fprint(os.Stdout, img); err != nil {
		log.Fatalf("failed to render image: %v", err)
	}

	rowsUsedByImg := imgTopPad + imgRows
	if rowsUsedByImg > 1 {
		fmt.Printf("\033[%dA", rowsUsedByImg-1)
	}
	fmt.Print("\r")

	for _, line := range tableLines {
		fmt.Printf("\033[%dC%s\n", tableOffsetCol, line)
	}

	if remaining := totalHeight - len(tableLines); remaining > 0 {
		fmt.Print(strings.Repeat("\n", remaining+1))
	}
}
