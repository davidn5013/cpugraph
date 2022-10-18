// package cpugraph plot a ascii grap of cpu usage
package cpugraph

import (
	"fmt"
	"math"
	"time"

	"github.com/guptarohit/asciigraph"
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
	for /*ever*/ {
		for i := 0; i < length; i++ {
			plot := asciigraph.Plot(data)
			fmt.Println(plot)
			data[i] = math.Floor(GetCpuUsage())
			asciigraph.Clear()
		}
		for i := 0; i < length; i++ {
			data[i] = 0
		}
	}
}
