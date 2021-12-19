package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Generates executable SQL from source, model, test, and analysis files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	compileCmd.Flags().Bool("full-refresh", false, "drop incremental models and fully-recalculate the incremental table")
	compileCmd.Flags().StringP("models", "m", "", "Specify the nodes to include.")
	compileCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	compileCmd.Flags().Bool("parse-only", false, "parse only")
	compileCmd.Flags().String("profile", "", "Which profile to load. Overrides setting in dbt_project.yml.")
	compileCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	compileCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	compileCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	compileCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	compileCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	compileCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	compileCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	compileCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(compileCmd)
}
