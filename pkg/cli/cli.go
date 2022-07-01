package cli

import "fmt"

type CLI struct {
	commandName string `json:"command_name"`
}

func New(options ...Option) (*CLI, error) {
	// Default options
	cli := &CLI{}

	for _, option := range options {
		if err := option(cli); err != nil {
			return nil, err
		}
	}

	return cli, nil
}

func (c *CLI) Run() error {
	fmt.Printf("welcome to %s!\n", c.commandName)

	return nil
}
