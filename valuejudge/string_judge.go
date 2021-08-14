package valuejudge

import "strings"

func IsStrBlank(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}
