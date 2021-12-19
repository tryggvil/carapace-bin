package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Execute snapshots defined in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().Bool("defer", false, "defer to the state variable for resolving unselected nodes")
	snapshotCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	snapshotCmd.Flags().StringP("models", "m", "", "Specify the nodes to include.")
	snapshotCmd.Flags().Bool("no-defer", false, "do not defer to the state variable for resolving unselected nodes")
	snapshotCmd.Flags().String("profile", "", "Which profile to load.")
	snapshotCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	snapshotCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	snapshotCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	snapshotCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	snapshotCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	snapshotCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	snapshotCmd.Flags().String("threads", "", "Specify number of threads to use while snapshotting tables.")
	snapshotCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(snapshotCmd)
}
