package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	nsjam := cli.NewApp()

	nsjam.Commands = []cli.Command{
		Version,
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
