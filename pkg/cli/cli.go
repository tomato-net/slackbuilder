package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CLI struct {
	commandName string `json:"command_name"`

	description string `json:"description"`

	cmd *cobra.Command `json:"cmd"`
}

func New(options ...Option) (*CLI, error) {
	// Default options
	cli := &CLI{}

	for _, option := range options {
		if err := option(cli); err != nil {
			return nil, err
		}
	}

	//cli.cmd = command.New(cli.commandName, cli.description)

	return cli, nil
}

func (c *CLI) Run() error {
	fmt.Printf("welcome to %s!\n", c.commandName)

	return nil
}
