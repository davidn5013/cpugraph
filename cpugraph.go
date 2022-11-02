// package cpugraph plot a ascii grap of cpu usage
package cpugraph

import (
	"math"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/pterm/pterm"
	"github.com/shirou/gopsutil/cpu"
)

// GetCpuUsage Procent usage of cpu in float64
func GetCpuUsage() (usedPercent float64) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic("Could not retrieve CPU usage details.")
	}
	usedPercent = cpuPercent[0]
	return usedPercent
}

// PlotCGrap Print a ascii graph of cpu usage for ever until ctrl-c
// length is the number measure point before graph start over
func PlotCGrap(length int) {
	data := make([]float64, length)

	// Turn off cursor
	pterm.Println("\033[?25l")

	// Using pterm Area to stopp the blink
	area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.

	for /*ever*/ {
		for i := 0; i < length; i++ {
			plot := asciigraph.Plot(data)
			area.Update(plot) // pterm update Area contents.
			data[i] = math.Floor(GetCpuUsage())
		}
		// for i := 0; i < length; i++ {
		// 	data[i] = 0
		// }
	}
}
