package commands

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
)

type ListNamespaces struct {
	Pid int `short:"p" long:"pid" description:"pid of process to list"`
}

func (command *ListNamespaces) Execute(args []string) error {
	process, _ := ps.FindProcess(command.Pid)
	if process == nil {
		return fmt.Errorf("no process found with pid: %d", command.Pid)
	}

	fmt.Println(process.Executable())
	return nil
}
