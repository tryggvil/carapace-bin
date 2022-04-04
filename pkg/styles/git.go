package styles

import "github.com/rsteube/carapace/pkg/style"

var Git = struct {
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

func init() {
	style.Register("git", &Git)
}
