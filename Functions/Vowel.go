package functions

import (
	"strings"
)

func spliter(str string) []string {
	var NewStr []string
	strs := ""
	for _, va := range str {
		if va == ' ' {
			strs += " "
			NewStr = append(NewStr, strings.TrimSpace(strs))
			strs = ""
		} else {
			strs += string(va)
		}
	}
	if len(strs) > 0 {
		NewStr = append(NewStr, strings.TrimSpace(strs))
		strs = ""
	}
	return NewStr
}

func Vowel(NewStr []string) string {
	temp := ""
	for i, va := range NewStr {
		if (va == "(" || va == ")") && i+1 < len(NewStr) && (NewStr[i+1] == "(" || NewStr[i+1] == ")" || NewStr[i+1] == "[" || NewStr[i+1] == "]") {
			temp += string(va)
		} else if (va == "[" || va == "]") && i+1 < len(NewStr) && (NewStr[i+1] == "(" || NewStr[i+1] == ")") {
			temp += string(va)
		} else {
			temp += string(va) + " "
		}
	}
	test := spliter(temp)
	NewStr = test
	laststr := ""

	for a := 0; a < len(NewStr); a++ {
		if NewStr[a] == "" || NewStr[a] == " " {
			continue
		}
		if len(NewStr[a]) == 2 {
			r := []rune(NewStr[a])
			if r[0] == 39 || r[0] == 40 || r[0] == 34 {
				if r[1] == 'A' && a+1 < len(NewStr) {
					if NewStr[a+1] == " " || NewStr[a+1] == "" {
						continue
					}
					runes := []rune(NewStr[a+1])
					next := runes[0]
					if next == 'a' || next == 'e' || next == 'i' || next == 'o' || next == 'u' || next == 'h' ||
						next == 'A' || next == 'E' || next == 'I' || next == 'O' || next == 'U' || next == 'H' {
						laststr += string(r[0]) + "An "
						continue
					}
				} else if r[1] == 'a' && a+1 < len(NewStr) {
					if NewStr[a+1] == " " || NewStr[a+1] == "" {
						continue
					}
					runes := []rune(NewStr[a+1])
					next := runes[0]
					if next == 'a' || next == 'e' || next == 'i' || next == 'o' || next == 'u' || next == 'h' ||
						next == 'A' || next == 'E' || next == 'I' || next == 'O' || next == 'U' || next == 'H' {
						laststr += string(r[0]) + "an "
						continue
					}
				}
			}
		}
		if NewStr[a] == "A" && a+1 < len(NewStr) {
			if NewStr[a+1] == " " || NewStr[a+1] == "" {
				continue
			}
			runes := []rune(NewStr[a+1])
			next := runes[0]
			if next == 'a' || next == 'e' || next == 'i' || next == 'o' || next == 'u' || next == 'h' ||
				next == 'A' || next == 'E' || next == 'I' || next == 'O' || next == 'U' || next == 'H' {
				laststr += "An "
				continue
			}
		} else if NewStr[a] == "a" && a+1 < len(NewStr) {
			if NewStr[a+1] == " " || NewStr[a+1] == "" {
				continue
			}
			runes := []rune(NewStr[a+1])
			next := runes[0]
			if next == 'a' || next == 'e' || next == 'i' || next == 'o' || next == 'u' || next == 'h' ||
				next == 'A' || next == 'E' || next == 'I' || next == 'O' || next == 'U' || next == 'H' {
				laststr += "an "
				continue
			}
		}
		// fmt.Println(NewStr[a])
		// if a+1 < len(NewStr) && NewStr[a+1] == ")" {
		// 	laststr += NewStr[a]
		// 	continue
		// }
		laststr += NewStr[a] + " "
	}

	return laststr
}
