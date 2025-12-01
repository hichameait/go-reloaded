package functions

import (
	"strconv"
	"strings"
	"unicode"
)

func RemoveFlags(s []string, i int) []string {
	if i < 0 || i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
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
						n = 0
					}
					value = strconv.Itoa(int(n))
				case "(bin)":
					n, err := strconv.ParseInt(value, 2, 64)
					if err != nil {
						n = 0
					}
					value = strconv.Itoa(int(n))
				case "(cap)":
					r := []rune(value)
					if len(r) > 0 {
						index := 0
						for i := 0; i < len(r); i++ {
							if unicode.IsLetter(r[i]) || unicode.IsDigit(r[i]) {
								index = i
								break
							}
						}
						r[index] = unicode.ToUpper(r[index])
						for dd := index + 1; dd < len(r); dd++ {
							r[dd] = unicode.ToLower(r[dd])
						}
						value = string(r)
					}
				case "(up)":
					r := []rune(value)
					if len(r) > 0 {
						for dd := 0; dd < len(r); dd++ {
							r[dd] = unicode.ToUpper(r[dd])
						}
						value = string(r)
					}
				case "(low)":
					r := []rune(value)
					if len(r) > 0 {
						for dd := 0; dd < len(r); dd++ {
							r[dd] = unicode.ToLower(r[dd])
						}
						value = string(r)
					}
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
