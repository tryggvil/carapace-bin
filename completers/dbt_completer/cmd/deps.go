package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Pull the most recent version of the dependencies listed in packages.yml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().String("profile", "", "Which profile to load dbt_project.yml")
	depsCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file")
	depsCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	depsCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	depsCmd.Flags().String("vars", "", "Supply variables to the project")
	rootCmd.AddCommand(depsCmd)
}
