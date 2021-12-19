package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Compile SQL and execute against the current target database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("defer", false, "defer to the state variable for resolving unselected nodes")
	runCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	runCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure.")
	runCmd.Flags().Bool("full-refresh", false, "dbt will drop incremental models and fully-recalculate the incremental table")
	runCmd.Flags().StringP("models", "m", "", "Specify the nodes to include.")
	runCmd.Flags().Bool("no-defer", false, "do not defer to the state variable for resolving unselected nodes")
	runCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	runCmd.Flags().String("profile", "", "Which profile to load.")
	runCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	runCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	runCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	runCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	runCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	runCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	runCmd.Flags().String("threads", "", "Specify number of threads to use while executing models.")
	runCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(runCmd)
}
