package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runOperationCmd = &cobra.Command{
	Use:   "run-operation",
	Short: "Run the named macro with any supplied arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runOperationCmd).Standalone()

	runOperationCmd.Flags().String("args", "", "Supply arguments to the macro")
	runOperationCmd.Flags().String("profile", "", "Which profile to load")
	runOperationCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file")
	runOperationCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	runOperationCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	runOperationCmd.Flags().String("vars", "", "Supply variables to the project")
	rootCmd.AddCommand(runOperationCmd)
}
