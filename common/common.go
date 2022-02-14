package common

import "strings"

func StringContains(raw string, elements []string) bool {
	for _, e := range elements {
		if strings.Contains(raw, e) {
			return true
		}
	}
	return false
}

func StringHasPrefix(raw string, elements []string) bool {
	for _, e := range elements {
		if strings.HasPrefix(raw, e) {
			return true
		}
	}
	return false
}

func StringHasSuffix(raw string, elements []string) bool {
	for _, e := range elements {
		if strings.HasSuffix(raw, e) {
			return true
		}
	}
	return false
}
