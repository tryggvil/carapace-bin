package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	aliasCmd.AddCommand(alias_deleteCmd)

	carapace.Gen(alias_deleteCmd).PositionalCompletion(
		action.ActionAliases(),
	)
}
