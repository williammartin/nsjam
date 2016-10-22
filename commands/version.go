package commands

import "fmt"

type Version struct{}

func (command *Version) Execute(args []string) error {
	fmt.Println("0.0.1")
	return nil
}
