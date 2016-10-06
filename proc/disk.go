package proc

import (
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/disk"
	"io"
	"os"
)

type DiskInfo struct {
	Properties []string
	Data       []string
	Writer     io.Writer
}

func NewDiskInfo() *DiskInfo {
	return &DiskInfo{
		Properties: []string{"Total", "Free", "Used", "UsedPercent"},
		Writer:     os.Stdout,
	}
}

func (self *DiskInfo) GetDiskInfo() {
	diskUsage, err := disk.DiskUsage("/")
	if err != nil {
		println(err.Error())
		return
	}
	table := tablewriter.NewWriter(self.Writer)
	usedPercent := humanize.Ftoa(diskUsage.UsedPercent) + "%"

	data := []string{
		humanize.Bytes(diskUsage.Total),
		humanize.Bytes(diskUsage.Free),
		humanize.Bytes(diskUsage.Used),
		usedPercent,
	}
	table.Append(data)
	table.SetHeader(self.Properties)
	table.Render() // Send output
}
