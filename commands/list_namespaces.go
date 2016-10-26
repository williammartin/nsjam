package commands

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
	"github.com/williammartin/nsodyssey"
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

	namespaces, _ := nsodyssey.Namespaces(command.Pid)

	for ns, inode := range namespaces {
		fmt.Printf("%s:%s\n", ns, inode)
	}

	return nil
}
