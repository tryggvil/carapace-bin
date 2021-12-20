package action

import (
	"io/ioutil"
	"os"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
)

type profile struct {
	Outputs map[string]struct {
		Type string
	}
}

func actionProfiles(f func(profiles map[string]profile) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support different path by flag value
		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := ioutil.ReadFile(home + "/.dbt/profiles.yml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var profiles map[string]profile
		if err := yaml.Unmarshal(content, &profiles); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(profiles)
	})
}

func ActionProfiles() carapace.Action {
	return actionProfiles(func(profiles map[string]profile) carapace.Action {
		vals := make([]string, 0)
		for profile := range profiles {
			vals = append(vals, profile)
		}
		return carapace.ActionValues(vals...)
	})
}
