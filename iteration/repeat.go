package iteration

import "strings"

func Repeat(ch string, x int) string {
	var repeated_ch strings.Builder
	for i := 0; i < x; i++ {
		repeated_ch.WriteString(ch)
	}
	return repeated_ch.String()
}
