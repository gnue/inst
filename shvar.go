package inst

import (
	"os"
	"regexp"
)

var (
	reVars = regexp.MustCompile(`\$(\w+|{(\w+)})`)
)

func EmptyEnvs(a []string) []string {
	r := make([]string, 0, len(a))

	for _, v := range a {
		if os.Getenv(v) == "" {
			r = append(r, v)
		}
	}

	return r
}

func EmptyVars(s string) []string {
	return EmptyEnvs(Vars(s))
}

func Vars(s string) []string {
	m := reVars.FindAllStringSubmatch(s, -1)
	a := make([]string, 0, len(m))

	for _, mm := range m {
		var v string

		if mm[2] == "" {
			v = mm[1]
		} else {
			v = mm[2]
		}

		a = append(a, v)
	}

	return uniq(a)
}

func uniq(a []string) []string {
	r := make([]string, 0, len(a))
	m := make(map[string]bool)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			r = append(r, v)
			m[v] = true
		}
	}

	return r
}
