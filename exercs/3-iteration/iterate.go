package iteration

import "strings"

func Repeat(char string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}