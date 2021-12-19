package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Load data from csv files into your data warehouse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(seedCmd).Standalone()

	seedCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	seedCmd.Flags().Bool("full-refresh", false, "Drop existing seed tables and recreate them")
	seedCmd.Flags().StringP("models", "m", "", "Specify the nodes to include.")
	seedCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	seedCmd.Flags().String("profile", "", "Which profile to load")
	seedCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	seedCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	seedCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	seedCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	seedCmd.Flags().Bool("show", false, "Show a sample of the loaded data in the terminal")
	seedCmd.Flags().String("state", "", "use the given directory as the source for json files to compare")
	seedCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	seedCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	seedCmd.Flags().String("vars", "", "Supply variables to the project")
	rootCmd.AddCommand(seedCmd)
}
