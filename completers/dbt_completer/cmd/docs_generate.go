package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docs_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docs_generateCmd).Standalone()

	docs_generateCmd.Flags().String("exclude", "", "Specify the models to exclude.")
	docs_generateCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	docs_generateCmd.Flags().StringP("models", "m", "", "Specify the nodes to include.")
	docs_generateCmd.Flags().Bool("no-compile", false, "Do not run \"dbt compile\" as part of docs generation")
	docs_generateCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml file")
	docs_generateCmd.Flags().String("profile", "", "Which profile to load.")
	docs_generateCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	docs_generateCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	docs_generateCmd.Flags().StringP("select", "s", "", "Specify the nodes to include.")
	docs_generateCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	docs_generateCmd.Flags().String("state", "", "use the given directory as the source for json files to compare with this project.")
	docs_generateCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	docs_generateCmd.Flags().String("threads", "", "Specify number of threads to use while executing models.")
	docs_generateCmd.Flags().String("vars", "", "Supply variables to the project.")
	docsCmd.AddCommand(docs_generateCmd)
}
