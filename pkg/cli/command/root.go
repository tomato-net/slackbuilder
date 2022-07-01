package command

import "github.com/spf13/cobra"

func NewRoot(name, description string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  name,
		Long: description,
	}

	return cmd
}
