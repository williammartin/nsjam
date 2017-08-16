package main

import (
	"errors"
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
		if ctx.Int("target") == 0 {
			return errors.New("Must pass a target pid")
		}

		process, err := ps.FindProcess(ctx.Int("target"))
		if process == nil && err == nil {
			return errors.New("Must pass a valid pid")
		}

		fmt.Println(process.Executable())

		pidNS, err := getPidNS(process.Pid())
		if err != nil {
			return err
		}

		fmt.Println(strings.TrimSpace(pidNS))
		return nil
	},
}

func getPidNS(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/ns/pid", pid))
}
