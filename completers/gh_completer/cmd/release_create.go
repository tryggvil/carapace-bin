package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_createCmd = &cobra.Command{
	Use:   "create <tag> [<files>...]",
	Short: "Create a new release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	release_createCmd.Flags().BoolP("draft", "d", false, "Save the release as a draft instead of publishing it")
	release_createCmd.Flags().StringP("notes", "n", "", "Release notes")
	release_createCmd.Flags().StringP("notes-file", "F", "", "Read release notes from `file`")
	release_createCmd.Flags().BoolP("prerelease", "p", false, "Mark the release as a prerelease")
	release_createCmd.Flags().String("target", "", "Target `branch` or full commit SHA (default: main branch)")
	release_createCmd.Flags().StringP("title", "t", "", "Release title")
	releaseCmd.AddCommand(release_createCmd)

	carapace.Gen(release_createCmd).FlagCompletion(carapace.ActionMap{
		"target": action.ActionBranches(release_createCmd),
	})

	carapace.Gen(release_createCmd).PositionalCompletion(
		action.ActionReleases(release_createCmd),
	)
	carapace.Gen(release_createCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
