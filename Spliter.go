package functions

import (
	"strings"
)

func check(str string, ind int) bool {
	for i := ind; i < len(str); i++ {
		if str[i] == '(' {
			return true
		}
		if str[i] == ')' {
			return false
		}
	}
	return false
}

func Spliter(str string) []string {
	textcopy := ""
	var myarr []string
	nextext := []rune(str)
	i := 0
	flage := false 

	for i = 0; i < len(nextext); i++ {
		ch := nextext[i]
		if ch == '('  && i-1 >=0 && nextext[i-1] == ' ' {
			cmd := ""
			k := i

			for k < len(nextext) {
				cmd += string(nextext[k])
				if nextext[k] == ')' {
					break
				}
				k++
			}
			if strings.HasPrefix(cmd, "(cap") || strings.HasPrefix(cmd, "(up") || strings.HasPrefix(cmd, "(low") {

				flage = false
				if strings.TrimSpace(textcopy) != "" {
					myarr = append(myarr, strings.TrimSpace(textcopy))
				}
				textcopy = ""
				myarr = append(myarr, strings.TrimSpace(cmd))
				i = k
				continue
			}

			flage = true
			if strings.TrimSpace(textcopy) != "" {
				myarr = append(myarr, strings.TrimSpace(textcopy))
			}
			textcopy = "("
			continue
		}

		if ch == ')' {
			if flage {
				textcopy += ")"
				myarr = append(myarr, strings.TrimSpace(textcopy))
				textcopy = ""
				flage = false
				continue
			}
		}
		
		if flage {
			if ch == ' ' {
				if strings.TrimSpace(textcopy) != "" {
					myarr = append(myarr, strings.TrimSpace(textcopy))
				}
				textcopy = ""
				continue
			}
			textcopy += string(ch)
			continue
		}

		if ch == ' ' {
			if strings.TrimSpace(textcopy) != "" {
				myarr = append(myarr, strings.TrimSpace(textcopy))
			}
			textcopy = ""
			continue
		}

		textcopy += string(ch)
	}

	if strings.TrimSpace(textcopy) != "" {
		myarr = append(myarr, strings.TrimSpace(textcopy))
	}
	return myarr

}

/// old spliter
// func Spliter(str string) []string {
// 	var out []string
// 	part := ""
// 	flg := true
// 	// str = CleanSymbols(str)
// 	// str = CleanQuotes(str)
// 	r := []rune(str)

// 	for i := 0; i < len(r); i++ {
// 		ch := r[i]
// 		if ch == '(' && i+1 <len(r) && r[i+1] != ' '{
// 			if strings.TrimSpace(part) != "" {
// 				out = append(out, strings.TrimSpace(part))
// 			}
// 			part = "("
// 			flg = false
// 			continue
// 		}

// 		if ch == ')' && i+1 < len(r) && r[i+1] == ' ' {
// 			part += ")"
// 			out = append(out, strings.TrimSpace(part))
// 			part = ""
// 			flg = true
// 			continue
// 		}

// 		if ch == ' ' && flg {
// 			if strings.TrimSpace(part) != "" {
// 				out = append(out, strings.TrimSpace(part))
// 			}
// 			part = ""
// 			continue
// 		}

// 		part += string(ch)
// 	}

// 	if strings.TrimSpace(part) != "" {
// 		out = append(out, strings.TrimSpace(part))
// 	}

// 	return out
// }
