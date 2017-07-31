package main

import (
	"fmt"
	"os"
	"strings"

	ps "github.com/mitchellh/go-ps"
	"github.com/urfave/cli"
)

func main() {
	nsjam := cli.NewApp()

	nsjam.Commands = []cli.Command{
		Version,
		ListNamespaces,
	}

	_ = nsjam.Run(os.Args)
}

var Version = cli.Command{
	Name: "version",
	Action: func(ctx *cli.Context) error {
		fmt.Println("0.0.1")

		return nil
	},
}

var ListNamespaces = cli.Command{
	Name: "list-namespaces",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "pid",
		},
		cli.IntFlag{
			Name: "target",
		},
	},
	Action: func(ctx *cli.Context) error {
		process, _ := ps.FindProcess(ctx.Int("target"))
		fmt.Println(process.Executable())

		pidNS, _ := getPidNS(process.Pid())
		fmt.Println(strings.TrimSpace(pidNS))
		return nil
	},
}

func getPidNS(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", pid))
}
