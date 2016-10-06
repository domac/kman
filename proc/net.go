package proc

import (
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
	"strconv"
)

type NetstatInfo struct {
	Properties []string
	Data       []string
	Writer     io.Writer
}

func NewNetstatInfo() *NetstatInfo {
	return &NetstatInfo{
		Properties: []string{"pid", "name", "port", "status"},
		Writer:     os.Stdout,
	}
}

func (self *NetstatInfo) GetNetstatInfo(port string) {
	niters, _ := net.NetConnections("all")
	table := tablewriter.NewWriter(self.Writer)
	for _, ns := range niters {
		pid := ns.Pid
		proc, _ := process.NewProcess(pid)
		name, _ := proc.Name()
		lport := strconv.Itoa(int(ns.Laddr.Port))
		status := ns.Status
		data := []string{strconv.Itoa(int(pid)), name, lport, status}

		if port == "all" {
			table.Append(data)
		} else {

			if lport == port {
				table.Append(data)
				break
			}

		}
	}
	table.SetHeader(self.Properties)
	table.Render() // Send output
}
