package proc

import (
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/mem"
	"io"
	"os"
)

type MemoryInfo struct {
	Properties []string
	Data       []string
	Writer     io.Writer
}

func NewMemoryInfo() *MemoryInfo {
	return &MemoryInfo{
		Properties: []string{"Total", "Available", "Free", "Used"},
		Data:       []string{"", "", "", "", ""},
		Writer:     os.Stdout,
	}
}

func (self *MemoryInfo) GetMemInfo() {

	v, _ := mem.VirtualMemory()
	self.Data = []string{humanize.Bytes(v.Total), humanize.Bytes(v.Available), humanize.Bytes(v.Free), humanize.Bytes(v.Used)}
	data := [][]string{
		self.Data,
	}
	table := tablewriter.NewWriter(self.Writer)
	table.SetHeader(self.Properties)
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
