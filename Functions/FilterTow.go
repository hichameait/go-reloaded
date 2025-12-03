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

func CapWord(words []string, count int) []string {
	start := len(words) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		if len(words[i]) > 0 {
			r := []rune(words[i])
			index := 0
			for j := 0; j < len(r); j++ {
				if unicode.IsLetter(r[j]) || unicode.IsDigit(r[j]) {
					index = j
					break
				}
			}
			r[index] = unicode.ToUpper(r[index])
			for j := index + 1; j < len(r); j++ {
				r[j] = unicode.ToLower(r[j])
			}
			words[i] = string(r)
		}
	}
	return words
}

func LowWord(words []string, count int) []string {
	start := len(words) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func UpWord(words []string, count int) []string {
	start := len(words) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

func FilterTow(NewStr []string) []string {
	for i := 0; i < len(NewStr); i++ {
		if strings.Contains(NewStr[i], "(cap,") && NewStr[i][0:5] == "(cap," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(cap, " + Num + ")"
			if NewStr[i] == str {
				words := []string{}
				for d := i - 1; d >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						w := strings.Fields(NewStr[d])
						words = append(w, words...)
					}
				}
				words = CapWord(words, num)
				NewStr[i] = ""
				wi := len(words) - 1
				for d := i - 1; d >= 0 && wi >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						old := strings.Fields(NewStr[d])
						for j := len(old) - 1; j >= 0 && wi >= 0; j-- {
							old[j] = words[wi]
							wi--
						}
						NewStr[d] = strings.Join(old, " ")
					}
				}
			}
		}
		if strings.Contains(NewStr[i], "(low,") && NewStr[i][0:5] == "(low," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(low, " + Num + ")"
			if NewStr[i] == str {
				words := []string{}
				for d := i - 1; d >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						w := strings.Fields(NewStr[d])
						words = append(w, words...)
					}
				}
				words = LowWord(words, num)
				NewStr[i] = ""
				wi := len(words) - 1
				for d := i - 1; d >= 0 && wi >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						old := strings.Fields(NewStr[d])
						for j := len(old) - 1; j >= 0 && wi >= 0; j-- {
							old[j] = words[wi]
							wi--
						}
						NewStr[d] = strings.Join(old, " ")
					}
				}
			}
		}
		if strings.Contains(NewStr[i], "(up,") && NewStr[i][0:4] == "(up," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			}
			str := "(up, " + Num + ")"
			if NewStr[i] == str {
				words := []string{}
				for d := i - 1; d >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						w := strings.Fields(NewStr[d])
						words = append(w, words...)
					}
				}
				words = UpWord(words, num)
				NewStr[i] = ""
				wi := len(words) - 1
				for d := i - 1; d >= 0 && wi >= 0; d-- {
					if len(NewStr[d]) > 0 && NewStr[d] != "" {
						old := strings.Fields(NewStr[d])
						for j := len(old) - 1; j >= 0 && wi >= 0; j-- {
							old[j] = words[wi]
							wi--
						}
						NewStr[d] = strings.Join(old, " ")
					}
				}
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
