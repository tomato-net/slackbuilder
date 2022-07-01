package command

import "github.com/spf13/cobra"

func New(name, description string) *cobra.Command {
	cmd := NewRoot(name, description)

	cmd.AddCommand(
		NewInit(),
		NewCreate(),
	)

	return cmd
}
