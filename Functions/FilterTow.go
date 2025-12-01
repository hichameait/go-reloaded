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

func FilterTow(NewStr []string) []string {
	for i := 0; i < len(NewStr); i++ {
		if strings.Contains(NewStr[i], "(cap,") && NewStr[i][0:5] == "(cap," {
			Num := ExtractNum(NewStr[i])
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			} else {
				if num >= len(NewStr[:i]) {
					num = len(NewStr[:i])
				}
				str := "(cap, " + Num + ")"
				if NewStr[i] == str && i > 0 {
					for d := 1; d <= num; d++ {
						if len(NewStr[i-d]) == 0 {
							continue
						}
						r := []rune(NewStr[i-d])
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
						NewStr[i-d] = string(r)
					}
					NewStr[i] = ""
				} else if NewStr[i] == str && i == 0 {
					NewStr[i] = ""
				}
			}

		}
		if strings.Contains(NewStr[i], "(low,") && NewStr[i][0:5] == "(low," {
			Num := ExtractNum(NewStr[i])
			str := "(low, " + Num + ")"
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			} else {
				if num >= len(NewStr[:i]) {
					num = len(NewStr[:i])
				}
				if NewStr[i] == str && i > 0 {
					for d := 1; d <= num; d++ {
						if len(NewStr[i-d]) == 0 {
							continue
						}
						NewStr[i-d] = strings.ToLower(NewStr[i-d])
					}
					NewStr[i] = ""
				} else if NewStr[i] == str && i == 0 {
					NewStr[i] = ""
				}
			}
		}
		if strings.Contains(NewStr[i], "(up,") && NewStr[i][0:4] == "(up," {
			Num := ExtractNum(NewStr[i])
			str := "(up, " + Num + ")"
			num, err := strconv.Atoi(Num)
			if err != nil {
				continue
			} else {
				if num >= len(NewStr[:i]) {
					num = len(NewStr[:i])
				}
				if NewStr[i] == str && i > 0 {
					for d := 1; d <= num; d++ {
						if len(NewStr[i-d]) == 0 {
							continue
						}
						NewStr[i-d] = strings.ToUpper(NewStr[i-d])
					}
					NewStr[i] = ""
				} else if NewStr[i] == str && i == 0 {
					NewStr[i] = ""
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
