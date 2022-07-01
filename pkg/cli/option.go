package cli

type Option func(*CLI) error

func WithCommandName(commandName string) Option {
	return func(cli *CLI) error {
		cli.commandName = commandName
		return nil
	}
}
