package pkg

import "strings"

func ClearLine(s string) string {
	return "\r" + strings.Repeat(" ", len(s)) + "\r"
}
