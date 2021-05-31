package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a new project or group variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	variable_setCmd.Flags().StringP("group", "g", "", "Set variable for a group")
	variable_setCmd.Flags().BoolP("masked", "m", false, "Whether the variable is masked")
	variable_setCmd.Flags().BoolP("protected", "p", false, "Whether the variable is protected")
	variable_setCmd.Flags().StringP("scope", "s", "*", "The environment_scope of the variable. All (*), or specific environments")
	variable_setCmd.Flags().StringP("type", "t", "env_var", "The type of a variable: {env_var|file}")
	variable_setCmd.Flags().StringP("value", "v", "", "The value of a variable")
	variableCmd.AddCommand(variable_setCmd)

	carapace.Gen(variable_setCmd).FlagCompletion(carapace.ActionMap{
		"type":  carapace.ActionValues("env_var", "file"),
		"group": action.ActionGroups(variableCmd),
		"scope": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			all := carapace.ActionValues("*").Invoke(c)
			environments := action.ActionEnvironments(variable_setCmd).Invoke(c)
			return all.Merge(environments).ToA()
		}),
	})

	carapace.Gen(variable_setCmd).PositionalCompletion(
		action.ActionVariables(variableCmd),
	)
}
