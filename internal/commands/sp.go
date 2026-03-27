package commands

import (
	"github.com/spf13/cobra"
)

func newSPCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sp",
		Short: "Manage service provider resources",
	}

	cmd.AddCommand(newSPResourceCommand())
	cmd.AddCommand(newSPProviderCommand())

	return cmd
}
