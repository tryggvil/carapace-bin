// package git contains git related actions
package git

var Style = struct {
	Branch     string
	Commit     string
	HeadCommit string
	Stash      string
	Tag        string
}{
	Branch:     "blue",
	Commit:     "default",
	HeadCommit: "bold",
	Stash:      "green",
	Tag:        "yellow",
}
