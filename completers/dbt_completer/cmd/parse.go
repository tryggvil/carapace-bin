package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parsed the project and provides information on performance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(parseCmd).Standalone()

	parseCmd.Flags().Bool("compile", false, "compile")
	parseCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	parseCmd.Flags().String("profile", "", "Which profile to load.")
	parseCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	parseCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")
	parseCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	parseCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	parseCmd.Flags().String("vars", "", "Supply variables to the project")
	parseCmd.Flags().Bool("write-manifest", false, "write manifest")
	rootCmd.AddCommand(parseCmd)
}
