package command

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Command struct {
	err error
}

func (c *Command) Run(name, command string) {
	if c.err != nil {
		return
	}

	var stdout bytes.Buffer
	cmd := exec.Command(command)
	cmd.Stdout = &stdout
	c.err = cmd.Run()
	if c.err != nil {
		fmt.Println("\x1b[31;1m[ERROR]\t%s\x1b[0m", name)
		return
	}

	fmt.Println("\x1b[32;1m[DONE]\t%s\x1b[0m", name)

	return
}

func (c *Command) Error() error {
	return c.err
}
