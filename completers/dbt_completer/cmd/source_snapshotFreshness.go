package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var source_snapshotFreshnessCmd = &cobra.Command{
	Use:   "snapshot-freshness",
	Short: "TODO",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_snapshotFreshnessCmd).Standalone()

	source_snapshotFreshnessCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	source_snapshotFreshnessCmd.Flags().StringP("output", "o", "", "Specify the output path for the json report")
	source_snapshotFreshnessCmd.Flags().String("profile", "", "Which profile to load.")
	source_snapshotFreshnessCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	source_snapshotFreshnessCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	source_snapshotFreshnessCmd.Flags().StringP("select", "s", "", "Specify the nodes to include")
	source_snapshotFreshnessCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	source_snapshotFreshnessCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare with this project.")
	source_snapshotFreshnessCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	source_snapshotFreshnessCmd.Flags().String("threads", "", "Specify number of threads to use.")
	source_snapshotFreshnessCmd.Flags().String("vars", "", "Supply variables to the project.")
	sourceCmd.AddCommand(source_snapshotFreshnessCmd)
}
