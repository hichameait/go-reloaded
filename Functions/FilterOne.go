package functions

import (
	"strconv"
	"strings"
	"unicode"
)

func Change(s string) []string {
	var str []string

	if s[0] == '(' && s[len(s)-1] == ')' {
		str = Split(s)
	}

	return str
}

func FilterOne(NewStr []string) []string {
	for i := 0; i < len(NewStr); i++ {
		flag := NewStr[i]
		if flag == "(hex)" || flag == "(bin)" || flag == "(up)" || flag == "(low)" || flag == "(cap)" {
			j := i - 1
			for j >= 0 && NewStr[j] == "" {
				j--
			}
			if j < 0 {
				NewStr[i] = ""
				continue
			}

			count := 0
			for k := i; k < len(NewStr) && NewStr[k] == flag; k++ {
				count++
			}

			value := NewStr[j]
			for k := 0; k < count; k++ {
				switch flag {
				case "(hex)":
					n, err := strconv.ParseInt(value, 16, 64)
					if err != nil {
						continue
					}
					value = strconv.Itoa(int(n))
				case "(bin)":
					n, err := strconv.ParseInt(value, 2, 64)
					if err != nil {
						continue
					}
					value = strconv.Itoa(int(n))
				case "(cap)":
					runes := []rune(value)
					part1 := ""
					part2 := value
					if len(runes) > 1 && runes[0] == '(' && runes[len(runes)-1] == ')' {
						inner := string(runes[1 : len(runes)-1])
						words := strings.Fields(inner)
						if len(words) > 1 {
							part1 = "(" + strings.Join(words[:len(words)-1], " ") + " "
							part2 = words[len(words)-1] + ")"
						}
					}
					r := []rune(part2)
					if len(r) > 0 {
						index := -1
						for i := 0; i < len(r); i++ {
							if unicode.IsLetter(r[i]) || unicode.IsDigit(r[i]) {
								index = i
								break
							}
						}
						if index != -1 {
							r[index] = unicode.ToUpper(r[index])
							for i := index + 1; i < len(r); i++ {
								r[i] = unicode.ToLower(r[i])
							}
						}
					}
					value = part1 + string(r)

				case "(up)":
					runes := []rune(value)
					part1 := ""
					part2 := value
					if len(runes) > 1 && runes[0] == '(' && runes[len(runes)-1] == ')' {
						inner := string(runes[1 : len(runes)-1])
						words := strings.Fields(inner)
						if len(words) > 1 {
							part1 = "(" + strings.Join(words[:len(words)-1], " ") + " "
							part2 = words[len(words)-1] + ")"
						}
					}
					r := []rune(part2)
					for i := 0; i < len(r); i++ {
						r[i] = unicode.ToUpper(r[i])
					}
					value = part1 + string(r)
				case "(low)":
					runes := []rune(value)
					part1 := ""
					part2 := value
					if len(runes) > 1 && runes[0] == '(' && runes[len(runes)-1] == ')' && strings.Contains(value, " ") {
						inner := string(runes[1 : len(runes)-1])
						words := strings.Fields(inner)

						if len(words) > 1 {
							part1 = "(" + strings.Join(words[:len(words)-1], " ") + " "
							part2 = words[len(words)-1] + ")"
						}
					}
					r := []rune(part2)
					for i := 0; i < len(r); i++ {
						r[i] = unicode.ToLower(r[i])
					}

					value = part1 + string(r)

				}
			}
			NewStr[j] = value
			for k := 0; k < count; k++ {
				NewStr[i+k] = ""
			}
			i += count - 1
		}

	}

	var LastArray []string
	for _, va := range NewStr {
		if len(va) != 0 && va != " " {
			LastArray = append(LastArray, strings.Trim(va, " "))
		} else {
			continue
		}
	}
	return LastArray
}
