package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the resources in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	listCmd.Flags().String("indirect-selection", "", "Select all tests that are adjacent to selected resources")
	listCmd.Flags().StringP("models", "m", "", "Specify the models to select and set the resource-type to 'model'.")
	listCmd.Flags().String("output", "", "set output ")
	listCmd.Flags().String("output-keys", "", "set output keys")
	listCmd.Flags().String("profile", "", "Which profile to load.")
	listCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	listCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	listCmd.Flags().String("resource-type", "", "set resource type")
	listCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	listCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	listCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	listCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	listCmd.Flags().String("vars", "", "Supply variables to the project.")
	rootCmd.AddCommand(listCmd)
}
