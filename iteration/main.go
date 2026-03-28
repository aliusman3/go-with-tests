package iteration

import "strings"

// Repeats a character n times and returns the resulting string
func Repeat(character string, n int) string {
	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
