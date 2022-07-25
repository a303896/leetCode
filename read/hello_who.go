package read

import (
	"fmt"
	"os"
	"strings"
)

func Hello() {
	if len(os.Args) > 1 {
		fmt.Println("Hello " + strings.Join(os.Args[1:], " ") + " !")
	}
}
