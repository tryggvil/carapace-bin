package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func Scan(path string, f func(scanner *bufio.Scanner) error) error {
	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return f(scanner)
}

func LintFilename(path string) error {
	r := regexp.MustCompile(`^	Use:   "(?P<cmd>[^ ]*)( [^"]+)?",`)

	return Scan(path, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			if r.MatchString(scanner.Text()) {
				cmd := normalize(r.FindStringSubmatch(scanner.Text())[1])
				if !strings.HasSuffix(filepath.Base(path), cmd+".go") {
					return fmt.Errorf("%v: file should has suffix '%v.go'", path, cmd)
				}
				return nil
			}
		}
		return nil
	})
}

func normalize(cmd string) string {
	normalized := make([]rune, 0)
	upper := false
	for _, r := range cmd {
		if r == '-' {
			upper = true
			continue
		}
		if upper {
			normalized = append(normalized, unicode.ToUpper(r))
			upper = false
		} else {
			normalized = append(normalized, r)
		}
	}
	return string(normalized)
}
