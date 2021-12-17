package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var source_freshnessCmd = &cobra.Command{
	Use:   "freshness",
	Short: "TODO freshness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_freshnessCmd).Standalone()

	source_freshnessCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	source_freshnessCmd.Flags().StringP("output", "o", "", "Specify the output path for the json report.")
	source_freshnessCmd.Flags().String("profile", "", "Which profile to load. Overrides setting in dbt_project.yml.")
	source_freshnessCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	source_freshnessCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	source_freshnessCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	source_freshnessCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	source_freshnessCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare with this project.")
	source_freshnessCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	source_freshnessCmd.Flags().String("threads", "", "Specify number of threads to use.")
	source_freshnessCmd.Flags().String("vars", "", "Supply variables to the project.")
	sourceCmd.AddCommand(source_freshnessCmd)
}
