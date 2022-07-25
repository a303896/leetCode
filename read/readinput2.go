package read

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input1 string
	err error
)

func Execute() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter something here:")
	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input content is:\n %s", input1)
	}
}

/**
read.Execute()

Enter something here:
Alpha
The input content is:
 Alpha
 */