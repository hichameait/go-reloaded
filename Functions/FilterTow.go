package functions

import (
	"strconv"
	"strings"
	"unicode"
)

func ExtractNum(str string) string {
	lastNum := ""
	flag := true
	for _, va := range str[5:] {
		if va != ')' && va != ' ' && flag {
			lastNum += string(va)
		} else if va == ')' {
			flag = false
		}
	}
	return lastNum
}

// getWordsFromParens splits "(word1 word2 word3)suffix" into words and returns suffix
func getWordsFromParens(word string) ([]string, string) {
	closeIdx := strings.LastIndex(word, ")")
	if closeIdx <= 0 {
		return nil, ""
	}
	inner := word[1:closeIdx]
	suffix := ""
	if closeIdx < len(word)-1 {
		suffix = word[closeIdx+1:]
	}
	return strings.Fields(inner), suffix
}

// rebuildParens rebuilds "(word1 word2 word3)suffix" from words and suffix
func rebuildParens(words []string, suffix string) string {
	return "(" + strings.Join(words, " ") + ")" + suffix
}

// capWord capitalizes a word
func capWord(word string) string {
	r := []rune(word)
	if len(r) == 0 {
		return word
	}
	index := 0
	for i := 0; i < len(r); i++ {
		if unicode.IsLetter(r[i]) || unicode.IsDigit(r[i]) {
			index = i
			break
		}
	}
	r[index] = unicode.ToUpper(r[index])
	for j := index + 1; j < len(r); j++ {
		r[j] = unicode.ToLower(r[j])
	}
	return string(r)
}

func FilterTow(NewStr []string) []string {
	for i := 0; i < len(NewStr); i++ {
		// Handle (cap, N)
		if len(NewStr[i]) >= 5 && NewStr[i][0:5] == "(cap," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(cap, " + Num + ")"
			if NewStr[i] == str && i > 0 {
				remaining := num
				// Go backwards and apply cap to 'num' words
				for j := i - 1; j >= 0 && remaining > 0; j-- {
					if NewStr[j] == "" {
						continue
					}
					// Check if word has parentheses like "(word1 word2)s"
					if len(NewStr[j]) > 2 && NewStr[j][0] == '(' && strings.Contains(NewStr[j], ")") {
						words, suffix := getWordsFromParens(NewStr[j])
						// Apply cap to words from the end
						for k := len(words) - 1; k >= 0 && remaining > 0; k-- {
							words[k] = capWord(words[k])
							remaining--
						}
						NewStr[j] = rebuildParens(words, suffix)
					} else {
						NewStr[j] = capWord(NewStr[j])
						remaining--
					}
				}
				NewStr[i] = ""
			} else if NewStr[i] == str && i == 0 {
				NewStr[i] = ""
			}
		}

		// Handle (low, N)
		if len(NewStr[i]) >= 5 && NewStr[i][0:5] == "(low," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(low, " + Num + ")"
			if NewStr[i] == str && i > 0 {
				remaining := num
				for j := i - 1; j >= 0 && remaining > 0; j-- {
					if NewStr[j] == "" {
						continue
					}
					if len(NewStr[j]) > 2 && NewStr[j][0] == '(' && strings.Contains(NewStr[j], ")") {
						words, suffix := getWordsFromParens(NewStr[j])
						for k := len(words) - 1; k >= 0 && remaining > 0; k-- {
							words[k] = strings.ToLower(words[k])
							remaining--
						}
						NewStr[j] = rebuildParens(words, suffix)
					} else {
						NewStr[j] = strings.ToLower(NewStr[j])
						remaining--
					}
				}
				NewStr[i] = ""
			} else if NewStr[i] == str && i == 0 {
				NewStr[i] = ""
			}
		}

		// Handle (up, N)
		if len(NewStr[i]) >= 4 && NewStr[i][0:4] == "(up," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(up, " + Num + ")"
			if NewStr[i] == str && i > 0 {
				remaining := num
				for j := i - 1; j >= 0 && remaining > 0; j-- {
					if NewStr[j] == "" {
						continue
					}
					if len(NewStr[j]) > 2 && NewStr[j][0] == '(' && strings.Contains(NewStr[j], ")") {
						words, suffix := getWordsFromParens(NewStr[j])
						for k := len(words) - 1; k >= 0 && remaining > 0; k-- {
							words[k] = strings.ToUpper(words[k])
							remaining--
						}
						NewStr[j] = rebuildParens(words, suffix)
					} else {
						NewStr[j] = strings.ToUpper(NewStr[j])
						remaining--
					}
				}
				NewStr[i] = ""
			} else if NewStr[i] == str && i == 0 {
				NewStr[i] = ""
			}
		}
	}

	var LastArray []string
	for _, va := range NewStr {
		if len(va) != 0 && va != " " {
			LastArray = append(LastArray, strings.Trim(va, " "))
		}
	}
	return LastArray
}
