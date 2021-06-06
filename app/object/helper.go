package object

import "strings"

// StringToBool string convert to bool
func StringToBool(check string) bool {
	switch strings.ToLower(check) {
	case "1":
		return true
	case "true":
		return true
	case "false":
		return false
	case "0":
		return false
	case "":
		return false
	}
	return false
}
