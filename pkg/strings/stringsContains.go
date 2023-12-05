package strings

import "strings"

func StringsContains(target string, sub []string) bool {
	for _, s := range sub {
		if strings.Contains(target, s) {
			return true
		}
	}
	return false
}
