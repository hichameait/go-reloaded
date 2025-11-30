package functions

import (
	"strings"
)




func check(str string, in int) bool {
	for i := in; i < len(str); i++ {
		if str[i] == '(' {
			return true
		}
		if str[i] == ')' {
			return false
		}
	}
	return false
}

func Spliter(str string) []string  {
	textcopy := ""
	var Str []string
	r := []rune(str)
	i := 0
	flage := false
	for i = 0; i < len(r)-1; i++ {

		switch r[i] {
		case '(':
			flage = true
		case ')':
			flage = false
		}
		if flage {
			if check(str, i) && r[i] == ' ' {
				if textcopy != "" {
					// Str = append(Str, textcopy)
					Str = append(Str, strings.TrimSpace(textcopy))
					textcopy = ""
				}
			}
		}

		if flage {
			textcopy += string(r[i])
			continue
		}
		if r[i] != '\n' && r[i] != ' ' {
			textcopy += string(r[i])
		}
		if r[i] == '\n' {
			textcopy += "\n"
		}

		if r[i] == ' ' && !flage {
			if textcopy != "" {
				// Str = append(Str, textcopy)
				Str = append(Str, strings.TrimSpace(textcopy))
				textcopy = ""
			}
		}
	}
	if i < len(r) {

		textcopy += string(r[i])
		if textcopy != "" {
			// Str = append(Str, textcopy)
			Str = append(Str, strings.TrimSpace(textcopy))
			textcopy = ""

		}
	}

	return Str
}
