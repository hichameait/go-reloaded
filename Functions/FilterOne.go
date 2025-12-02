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

// applyToLastWordInParens applies transformation to the last word inside parenthesized content
// For example: "(test1 test2 test3)" with cap -> "(test1 test2 Test3)"
func applyToLastWordInParens(value string, transform func(string) string) string {
	if len(value) > 2 && value[0] == '(' && value[len(value)-1] == ')' {
		inner := value[1 : len(value)-1]
		parts := strings.Fields(inner)
		if len(parts) > 0 {
			parts[len(parts)-1] = transform(parts[len(parts)-1])
			return "(" + strings.Join(parts, " ") + ")"
		}
	}
	return transform(value)
}

func capitalizeWord(word string) string {
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
	for dd := index + 1; dd < len(r); dd++ {
		r[dd] = unicode.ToLower(r[dd])
	}
	return string(r)
}

func uppercaseWord(word string) string {
	r := []rune(word)
	for dd := 0; dd < len(r); dd++ {
		r[dd] = unicode.ToUpper(r[dd])
	}
	return string(r)
}

func lowercaseWord(word string) string {
	r := []rune(word)
	for dd := 0; dd < len(r); dd++ {
		r[dd] = unicode.ToLower(r[dd])
	}
	return string(r)
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
					// For hex, we don't handle parenthesized groups specially
					n, err := strconv.ParseInt(value, 16, 64)
					if err != nil {
						continue
					}
					value = strconv.Itoa(int(n))
				case "(bin)":
					// For bin, we don't handle parenthesized groups specially
					n, err := strconv.ParseInt(value, 2, 64)
					if err != nil {
						continue
					}
					value = strconv.Itoa(int(n))
				case "(cap)":
					value = applyToLastWordInParens(value, capitalizeWord)
				case "(up)":
					value = applyToLastWordInParens(value, uppercaseWord)
				case "(low)":
					value = applyToLastWordInParens(value, lowercaseWord)
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
		}
	}
	return LastArray
}
