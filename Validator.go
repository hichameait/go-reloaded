package functions

import (
	"fmt"
	"os"
	"strings"
)

func Isvalid() bool {
	if len(os.Args) != 3 {
		fmt.Println("Error in Args fix it :)")
		return false
	}

	if !(strings.Contains(os.Args[1], ".txt") && strings.Contains(os.Args[2], ".txt")) {
		fmt.Println("File Extension not allowed")
		return false
	}
	return true
}
