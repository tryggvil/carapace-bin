package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List and filter issues in this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issue_listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	issue_listCmd.Flags().StringSliceP("label", "l", nil, "Filter by labels")
	issue_listCmd.Flags().StringP("state", "s", "open", "Filter by state: {open|closed|all}")
	issue_listCmd.Flags().BoolP("web", "w", false, "Open the browser to list the issue(s)")
	issue_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of issues to fetch")
	issue_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	issue_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_listCmd.Flags().String("mention", "", "Filter by mention")
	issue_listCmd.Flags().StringP("milestone", "m", "", "Filter by milestone `number` or `title`")
	issue_listCmd.Flags().StringP("search", "S", "", "Search issues with `query`")
	issue_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")

	issueCmd.AddCommand(issue_listCmd)

	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionAssignableUsers(issue_listCmd),
		"author":   action.ActionUsers(issue_listCmd, action.UserOpts{Users: true}),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"mention":   action.ActionAssignableUsers(issue_listCmd),
		"milestone": action.ActionMilestones(issue_listCmd),
		"state":     carapace.ActionValues("open", "closed", "all"),
	})
}
