package functions

import (
	"strings"
)

func Lines(line string) string {
	words := Spliter(line)
	words = FilterOne(words)
	words = FilterTow(words)
	last := Vowel(words)
	last = CleanSymbols(last)
	last = CleanQuotes(last)
	if strings.Contains(last, " , ") {
		last = CleanSymbols(last)
		last = CleanQuotes(last)
	}
	return last
}
