package main

import (
	"github.com/codegangsta/cli"
	"os"
	"study2016/kman/proc"
)

//应用启动Action
func appAction(c *cli.Context) (err error) {

	memFlag := c.Bool("mem")
	cpuFlag := c.Bool("cpu")
	diskFlag := c.Bool("disk")
	kname := c.String("pname")
	kport := c.String("port")

	if memFlag {
		//列出内存的详细信息
		meminfo := proc.NewMemoryInfo()
		meminfo.GetMemInfo()
	}
	if cpuFlag {
		//列出cpu的详细信息
		cpuinfo := proc.NewCpuInfo()
		cpuinfo.GetCpuInfo()
	}

	if diskFlag {
		diskinfo := proc.NewDiskInfo()
		diskinfo.GetDiskInfo()
	}

	if kname != "" {
		//执行根据进程名称删除
		prcessinfo := proc.NewProcessInfo()
		prcessinfo.GetProcessInfo(kname)
	}

	if kport != "" {
		//删除指定端口所在的进程
		netstatinfo := proc.NewNetstatInfo()
		netstatinfo.GetNetstatInfo(kport)
	}
	return nil
}

func main() {
	proc.FlagsInit()
	app := cli.NewApp()
	app.Name = "kman"
	app.Usage = "useful for getting infomation in our system"
	app.Version = proc.APP_VERSION
	app.Flags = proc.GetAppFlags()
	app.Action = proc.ActionWrapper(appAction)
	app.Run(os.Args)
}
