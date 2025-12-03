package functions

import (
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

func CleanQuotes(s string) string {
	runes := []rune(s)
	sb := ""
	seps := ".!,?;:"
	ln := len(runes)

	for i := 0; i < ln; i++ {
		r := runes[i]

		if r != '\'' {
			sb += string(r)
			continue
		}
		if i > 0 && i < ln-1 && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
			sb += "'"
			continue
		}

		j := i + 1
		for j < ln {
			if runes[j] != '\'' {
				j++
				continue
			}
			if j > i && j < ln-1 && IsAlpha(runes[j-1]) && IsAlpha(runes[j+1]) {
				j++
				continue
			}
			break
		}
		if j >= ln || runes[j] != '\'' {
			sb += "'"
			continue
		}

		inner := strings.TrimSpace(string(runes[i+1 : j]))

		if len(sb) > 0 && !strings.HasSuffix(sb, " ") {
			sb += " "
		}
		sb += "'" + inner + "'"

		if j+1 < ln && runes[j+1] != ' ' && !strings.ContainsRune(seps, runes[j+1]) {
			sb += " "
		}
		i = j
	}

	mts := regexp.MustCompile(`[ ]{2,}`)
	clean := mts.ReplaceAllString(sb, " ")

	// punct fix
	out := ""
	inQ := false
	rn := []rune(clean)
	for i := 0; i < len(rn); i++ {
		if rn[i] == '\'' {
			inQ = !inQ
			out += string(rn[i])
			continue
		}
		if !inQ && rn[i] == ' ' && i+1 < len(rn) && strings.ContainsRune(seps, rn[i+1]) {
			continue
		}
		out += string(rn[i])
	}

	return strings.TrimSpace(out)
}
