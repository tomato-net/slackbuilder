package command

import "github.com/spf13/cobra"

func NewInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialise the project",
		Long:  "Initialise the project",
	}

	return cmd
}
