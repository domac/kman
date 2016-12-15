package proc

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"os/exec"
	"strings"
)

type TimeWaitInfo struct {
	Properties []string
	Data       [][]string
	Writer     io.Writer
}

func NewTimeWaitInfo() *TimeWaitInfo {
	return &TimeWaitInfo{Writer: os.Stdout}
}

type SimpleWriter struct {
	Cache []byte
}

func (s *SimpleWriter) Write(p []byte) (n int, err error) {
	s.Cache = p
	return len(p), nil
}

func NewSimpleWriter() *SimpleWriter {
	return &SimpleWriter{
		Cache: make([]byte, 1024)}
}

//netstat -ant | awk '/^tcp/ {++S[$NF] } END {for(a in S) print a, S[a]}'
//netstat -ant | awk '/^tcp/ {++S[$NF] } END {for(a in S) print a, S[a], "#"}'
func (self *TimeWaitInfo) GetTimeWaitInfo() {
	cmd := exec.Command("netstat", "-ant")
	awkCmd := exec.Command("awk", "/^tcp/ {++S[$NF] } END {for(a in S) print a, S[a]}")

	sw := NewSimpleWriter()

	awkCmd.Stdin, _ = cmd.StdoutPipe()
	awkCmd.Stdout = sw

	_ = awkCmd.Start()
	_ = cmd.Run()
	_ = awkCmd.Wait()

	var b2 bytes.Buffer
	awkCmd.Stdout = &b2

	buf := []byte{}
	awkCmd.Stdout.Write(buf)

	//结果缓存
	cache := sw.Cache
	cacheList := strings.Split(string(cache), "\n")

	titles := []string{}
	row := []string{}

	for _, line := range cacheList {
		data := strings.Split(line, " ")
		if len(data) < 2 {
			continue
		}

		titles = append(titles, data[0])
		row = append(row, data[1])
	}

	table := tablewriter.NewWriter(self.Writer)
	table.SetHeader(titles)
	table.Append(row)
	table.Render()
}
