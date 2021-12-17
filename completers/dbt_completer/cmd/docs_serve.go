package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docs_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docs_serveCmd).Standalone()

	docs_serveCmd.Flags().Bool("no-browser", false, "")
	docs_serveCmd.Flags().String("port", "", "Specify the port number for the docs server.")
	docs_serveCmd.Flags().String("profile", "", "Which profile to load. Overrides setting in dbt_project.yml.")
	docs_serveCmd.Flags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file.")
	docs_serveCmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file.")
	docs_serveCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	docs_serveCmd.Flags().String("vars", "", "Supply variables to the project.")
	docsCmd.AddCommand(docs_serveCmd)
}
