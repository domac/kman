package proc

import (
	"github.com/codegangsta/cli"
)

//参数初始化
func FlagsInit() {

	AddFlagBool(cli.BoolFlag{
		Name:  "mem",
		Usage: "display the system memory info",
	})

	AddFlagBool(cli.BoolFlag{
		Name:  "cpu",
		Usage: "display the system cpu info",
	})

	AddFlagBool(cli.BoolFlag{
		Name:  "disk",
		Usage: "display the system disk info",
	})

	AddFlagString(cli.StringFlag{
		Name:  "pname",
		Usage: "find the process with name",
	})

	AddFlagString(cli.StringFlag{
		Name:  "port",
		Usage: "find the process with portl; eg: --port 8080 or --port all",
	})
}
