package styles

import "github.com/rsteube/carapace/pkg/style"

var Gh = struct {
	Draft         string
	JobFailed     string
	JobInProgress string
	JobSuccess    string
	StateClosed   string
	StateMerged   string
	StateOpen     string
}{
	Draft:         style.Gray,
	JobFailed:     style.Red,
	JobInProgress: style.Yellow,
	JobSuccess:    style.Green,
	StateClosed:   style.Red,
	StateMerged:   style.Magenta,
	StateOpen:     style.Green,
}

func init() {
	style.Register("gh", &Gh)
}
