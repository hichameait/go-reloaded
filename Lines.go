package functions

import (
	// "fmt"
	"strings"
)

func Lines(line string) string {
	words := Spliter(line)
	words = FilterOne(words)
	words = FilterTow(words)
	last := Vowel(words)
	last = CleanSymbols(last)
	last = CleanQuotes(last)
	// fmt.Println(last)
	if strings.Contains(last, " , ") {
		last = CleanSymbols(last)
		last = CleanQuotes(last)
	}
	return last
}
