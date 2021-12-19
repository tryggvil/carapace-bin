package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs tests on data in deployed models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("defer", false, "defer to the state variable for resolving unselected nodes")
	testCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	testCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first test failure.")
	testCmd.Flags().String("indirect-selection", "", "Select all tests that are adjacent to selected resources")
	testCmd.Flags().StringP("models", "m", "", "Specify the nodes to include")
	testCmd.Flags().Bool("no-defer", false, "do not defer to the state variable for resolving unselected nodes")
	testCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	testCmd.Flags().String("profile", "", "Which profile to load")
	testCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	testCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	testCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	testCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	testCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	testCmd.Flags().Bool("store-failures", false, "Store test results (failing rows) in the database")
	testCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	testCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	testCmd.Flags().String("vars", "", "Supply variables to the project")
	rootCmd.AddCommand(testCmd)
}
