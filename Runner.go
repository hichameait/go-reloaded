package functions

import (
	"fmt"
	"os"
	"strings"
)

func Runner() {
	if !Isvalid() {
		return
	}
	containt, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(string(containt)) == 0 {
		fmt.Println("The based file is empty")
		return
	}

	str := string(containt)
	Line := ""
	var Strarray []string
	for i, r := range str {
		char := string(r)
		if char == "\n" {
			processed := Lines(Line)
			// if processed == "error" {
			// 	return
			// }
			Strarray = append(Strarray, processed)
			Line = ""
			continue
		}
		if i == len(str)-1 {
			Line += char
			processed := Lines(Line)
			Strarray = append(Strarray, processed)
			Line = ""
			continue
		}
		Line += char
	}
	Line = ""
	for index, va := range Strarray {
		if index != len(Strarray)-1{
			Line += strings.TrimSpace(va) + "\n"
		}else {
			Line += strings.TrimSpace(va)
		}
	}

	MakeFile(Line)
}
