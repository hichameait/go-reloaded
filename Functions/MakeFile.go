package functions

import (
	"fmt"
	"os"
)

func MakeFile(laststr string) {
	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error to Create file")
		return
	}
	_, err = file.WriteString(laststr)
	if err != nil {
		fmt.Println("Error to Write the string")
		return
	}
}
