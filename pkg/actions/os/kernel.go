package os

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
)

func currentRelease(c carapace.Context) (string, error) {
	if output, err := c.Command("uname", "-r").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}

// ActionKernelModulesLoaded completes currently loaded kernel modules
//   ac97_bus
//   crc32c_intel
func ActionKernelModulesLoaded() carapace.Action {
	return carapace.ActionExecCommand("lsmod")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			vals = append(vals, strings.Fields(line)[0])
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionKernelModules completes kernel modules
//   ac97_bus
//   crc32c_intel
func ActionKernelModules(basedir string, release string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.TrimSpace(release) == "" {
			var err error
			if release, err = currentRelease(c); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		vals := make([]string, 0)
		err := filepath.WalkDir(fmt.Sprintf("%v/lib/modules/%v/kernel", basedir, release), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				name := filepath.Base(path)
				name = strings.TrimSuffix(name, ".xz")
				name = strings.TrimSuffix(name, ".ko")
				vals = append(vals, name)
			}
			return nil
		})
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionKernelReleases completes kernel releases
//   5.10.79-1-MANJARO
//   5.4.159-1-MANJARO
func ActionKernelReleases(basedir string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		files, err := ioutil.ReadDir(basedir + "/lib/modules")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if file.IsDir() {
				vals = append(vals, file.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
