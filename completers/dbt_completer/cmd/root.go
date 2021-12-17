package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dbt",
	Short: "An ELT tool for managing your SQL transformations and data models",
	Long:  "https://www.getdbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Display debug logging during dbt execution.")
	rootCmd.Flags().String("event-buffer-size", "", "Sets the max number of events to buffer in EVENT_HISTORY")
	rootCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure.")
	rootCmd.Flags().String("log-format", "", "Specify the log format, overriding the command's default.")
	rootCmd.Flags().Bool("no-anonymous-usage-stats", false, "Do not send anonymous usage stat to dbt Labs")
	rootCmd.Flags().Bool("no-partial-parse", false, "Disallow partial parsing.")
	rootCmd.Flags().Bool("no-static-parser", false, "Disables the static parser.")
	rootCmd.Flags().Bool("no-use-colors", false, "Do not colorize the output DBT prints to the terminal.")
	rootCmd.Flags().Bool("no-version-check", false, "If set, skip ensuring dbt's version matches the one specified in the dbt_project.yml file")
	rootCmd.Flags().Bool("no-write-json", false, "If set, skip writing the manifest and run_results.json files to disk")
	rootCmd.Flags().Bool("partial-parse", false, "Allow for partial parsing by looking for and writing to a pickle file in the target directory.")
	rootCmd.Flags().String("printer-width", "", "Sets the width of terminal output")
	rootCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	rootCmd.Flags().StringP("record-timing-info", "r", "", "When this option is passed, dbt will output low-level timing stats to the specified file.")
	rootCmd.Flags().Bool("use-colors", false, "Colorize the output DBT prints to the terminal.")
	rootCmd.Flags().Bool("use-experimental-parser", false, "Enables experimental parsing features.")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("warn-error", false, "If dbt would normally warn, instead raise an exception.")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "show this help message and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"record-timing-info": carapace.ActionFiles(),
		"log-format":         carapace.ActionValues("text", "json", "default"),
		"profiles-dir":       carapace.ActionDirectories(),
	})
}
