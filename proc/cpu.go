package proc

import (
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/cpu"
	"io"
	"os"
)

type CpuInfo struct {
	Properties []string
	Data       [][]string
	Writer     io.Writer
}

func NewCpuInfo() *CpuInfo {
	return &CpuInfo{
		Properties: []string{"cpu", "user", "system", "idle", "idle percentage", "nice", "iowait"},
		Writer:     os.Stdout,
	}
}

func (self *CpuInfo) GetCpuInfo() {
	cpuTimes, err := cpu.CPUTimes(true)

	if nil != err {
		return
	}

	table := tablewriter.NewWriter(self.Writer)
	table.SetHeader(self.Properties)

	for _, cpuTime := range cpuTimes {
		cpu := cpuTime.CPU
		user := cpuTime.User
		system := cpuTime.System
		idle := cpuTime.Idle
		nice := cpuTime.Nice
		iowait := cpuTime.Iowait

		idlePercent := (idle / (idle + user + system)) * 100

		ip := humanize.Ftoa(idlePercent) + "%"

		data := []string{
			cpu,
			humanize.Ftoa(user),
			humanize.Ftoa(system),
			humanize.Ftoa(idle),
			ip,
			humanize.Ftoa(nice),
			humanize.Ftoa(iowait),
		}
		table.Append(data)
	}

	table.Render()
}
