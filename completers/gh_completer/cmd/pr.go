package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr <command>",
	Short: "Manage pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	prCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(prCmd)

	carapace.Gen(prCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionOwnerRepositories(prCmd),
	})
}
