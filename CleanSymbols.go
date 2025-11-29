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

// func CleanQuotes(laststr string) string {

//     runes := []rune(laststr)
//     temp := ""

//     for i := 0; i < len(runes); i++ {
//         ch := runes[i]
//         if ch != '\'' {
//             temp += string(ch)
//             continue
//         }
//         if i+1 < len(runes) && runes[i+1] != ' ' {
//             temp += string(ch)
//         } else if i-1 >= 0 && runes[i-1] != ' ' {
//             temp += string(ch)
//         } else {
//             temp += string(ch)
//             i++
//         }
//     }
//     re3 := regexp.MustCompile(`[ ]{2,}`)
//     return re3.ReplaceAllString(temp, " ")

// }

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func CleanQuotes(s string) string {
	runes := []rune(s)
	sb := ""
	n := len(runes)
	seps := ".!,?;:"
	for i := 0; i < n; i++ {
		r := runes[i]
		if r != '\'' {
			sb += string(r)
			continue
		}
		if i > 0 && i < n-1 && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
			sb += "'"
			continue
		}
		j := i + 1
		for j < n && runes[j] != '\'' {
			j++
		}
		if j == n {
			sb += "'"
			continue
		}

		inner := strings.TrimSpace(string(runes[i+1 : j]))
		sb += " '" + inner + "' "
		i = j
	}
	
	re3 := regexp.MustCompile(`[ ]{2,}`)
	str := []rune(re3.ReplaceAllString(sb, " "))
	

	output := ""
	inQuote := false
	for i := 0; i < len(str); i++ {
		if str[i] == '\'' {
			inQuote = !inQuote
			output += string(str[i])
			continue
		}
		if !inQuote &&
			len(output) != 0 &&
			str[i] == ' ' &&
			i+1 < len(str) &&
			strings.Contains(seps, string(str[i+1])) {
			continue
		}else{
			output += string(str[i])
		}
	}
	return output
}
