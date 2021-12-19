package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run all Seeds, Models, Snapshots, and tests in DAG order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("defer", false, "defer to the state variable for resolving unselected nodes")
	buildCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	buildCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure.")
	buildCmd.Flags().Bool("full-refresh", false, "drop incremental models and fully-recalculate the incremental table")
	buildCmd.Flags().String("indirect-selection", "", "Select all tests that are adjacent to selected resources")
	buildCmd.Flags().Bool("no-defer", false, "do not defer to the state variable for resolving unselected nodes")
	buildCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	buildCmd.Flags().String("profile", "", "Which profile to load.")
	buildCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	buildCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	buildCmd.Flags().String("resource-type", "", "set resource type")
	buildCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	buildCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	buildCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	buildCmd.Flags().Bool("store-failures", false, "Store test results (failing rows) in the database")
	buildCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	buildCmd.Flags().String("threads", "", "Specify number of threads to use while executing models.")
	buildCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(buildCmd)
}
