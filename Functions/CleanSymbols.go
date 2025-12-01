package functions

import (
	// "fmt"
	"regexp"
	"strings"
	"unicode"
)

func CleanSymbols(laststr string) string {
	seps := ".!,?;:"
	str := ""
	runes := []rune(laststr)
	for i := 0; i < len(runes); i++ {
		ch := string(runes[i])
		if strings.Contains(seps, ch) {
			seq := ch
			j := i + 1
			for j < len(runes) && strings.Contains(seps, string(runes[j])) {
				seq += string(runes[j])
				j++
			}
			str = strings.TrimRight(str, " ")
			str += seq
			if j < len(runes) && !strings.Contains(seps, string(runes[j])) && runes[j] != ' ' {
				str += " "
			}
			i = j - 1
			continue
		}
		str += ch
	}
	return str
}

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

var reMultiSpace = regexp.MustCompile(`[ ]{2,}`)

func CleanQuotes(s string) string {
	runes := []rune(s)
	var sb strings.Builder
	n := len(runes)
	seps := ".!,?;:"

	for i := 0; i < n; i++ {
		r := runes[i]
		if r != '\'' {
			sb.WriteRune(r)
			continue
		}
		if i > 0 && i < n-1 && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
			sb.WriteRune('\'')
			continue
		}
		j := i + 1
		for j < n && runes[j] != '\'' {
			j++
		}
		if j == n {
			sb.WriteRune('\'')
			continue
		}

		inner := strings.TrimSpace(string(runes[i+1 : j]))
		if sb.Len() > 0 && !strings.HasSuffix(sb.String(), " ") {
			sb.WriteRune(' ')
		}
		sb.WriteString("'" + inner + "'")
		if j+1 < n && runes[j+1] != ' ' && !strings.ContainsRune(seps, runes[j+1]) {
			sb.WriteRune(' ')
		}
		i = j
	}

	str := []rune(reMultiSpace.ReplaceAllString(sb.String(), " "))

	var output strings.Builder
	inQuote := false
	for i := 0; i < len(str); i++ {
		if str[i] == '\'' {
			inQuote = !inQuote
			output.WriteRune(str[i])
			continue
		}
		if !inQuote &&
			output.Len() != 0 &&
			str[i] == ' ' &&
			i+1 < len(str) &&
			strings.ContainsRune(seps, str[i+1]) {
			continue
		}
		output.WriteRune(str[i])
	}

	return strings.TrimSpace(output.String())
}
