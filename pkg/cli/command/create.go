package command

import (
	"github.com/spf13/cobra"
)

func NewCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "create",
		SuggestFor: []string{"new"},
		Short:      "Scaffold a slack command",
		Long:       "Scaffold a slack command",
	}

	cmd.AddCommand(
		NewCreateSlashCMD(),
	)

	return cmd
}

func NewCreateSlashCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slashcmd",
		Short: "Scaffold a slack slash command",
		Long:  "Scaffold a slack slash command",
	}

	return cmd
}
