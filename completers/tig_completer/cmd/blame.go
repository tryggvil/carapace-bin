package cmd

import (
	"github.com/rsteube/carapace"
	git "github.com/rsteube/carapace-bin/completers/git_completer/cmd"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:                "blame",
	Short:              "Show what revision and author last modified each line of a file",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"blame"}, c.Args...)
			return git.ActionExecute().Chdir(blameCmd.Root().Flag("C").Value.String()).Invoke(c).ToA()
		}),
	)
}
