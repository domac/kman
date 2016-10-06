package proc

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
)

var appFlags = map[string]cli.Flag{}

func ActionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := action(context); err != nil {
			log.Println(err.Error())
		}
	}
}

//获取应用入参
func GetAppFlags() (afs []cli.Flag) {
	for _, f := range appFlags {
		afs = append(afs, f)
	}
	return
}

func AddFlagString(sf cli.StringFlag) cli.StringFlag {
	if _, ok := appFlags[sf.Name]; ok {
		panic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}

func AddFlagBool(sf cli.BoolFlag) cli.BoolFlag {
	if _, ok := appFlags[sf.Name]; ok {
		panic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}

func AddFlagInt(sf cli.IntFlag) cli.IntFlag {
	if _, ok := appFlags[sf.Name]; ok {
		panic(fmt.Sprintf("flag %s denined", sf.Name))
	} else {
		appFlags[sf.Name] = sf
	}
	return sf
}
