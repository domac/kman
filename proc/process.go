package proc

import (
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
	"strconv"
	"strings"
)

type ProcessInfo struct {
	Properties []string
	Data       []string
	Writer     io.Writer
}

func NewProcessInfo() *ProcessInfo {
	return &ProcessInfo{
		Properties: []string{"pid", "name", "parent"},
		Writer:     os.Stdout,
	}
}

func (self *ProcessInfo) GetProcessInfo(name string) {

	pids, err := process.Pids()
	if err != nil {
		return
	}

	table := tablewriter.NewWriter(self.Writer)
	table.SetHeader(self.Properties)
	for _, pid := range pids {
		proc, _ := process.NewProcess(pid)
		pname := getStringVal(proc.Name())

		t_pname := strings.ToUpper(pname)
		t_name := strings.ToUpper(name)

		if !strings.Contains(t_pname, t_name) {
			continue
		}

		ppid := getInt32Val(proc.Ppid())
		data := []string{strconv.Itoa(int(pid)), pname, ppid}
		table.Append(data)
	}

	table.Render() // Send output
}

func getStringVal(val string, err error) string {
	if err != nil {
		return ""
	}
	return val
}

func getInt32Val(val int32, err error) string {
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(val))
}
