package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflow_disableCmd = &cobra.Command{
	Use:   "disable [<workflow-id> | <workflow-name>]",
	Short: "Disable a workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	workflowCmd.AddCommand(workflow_disableCmd)

	carapace.Gen(workflow_disableCmd).PositionalCompletion(
		action.ActionWorkflows(workflow_disableCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	)
}
