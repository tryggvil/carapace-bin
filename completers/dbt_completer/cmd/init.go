package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new DBT project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("profile", "", "Which profile to load.")
	initCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	initCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	initCmd.Flags().BoolP("skip-profile-setup", "s", false, "Skip interative profile setup.")
	initCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	initCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(initCmd)
}
