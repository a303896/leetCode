package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ExecRemove() {
	input,_ := os.Open(DIR + "products.txt")
	output,_ := os.OpenFile(DIR + "remove_3till5char.txt", os.O_WRONLY | os.O_CREATE, 0666)
	defer input.Close()
	defer output.Close()
	inputReader := bufio.NewReader(input)
	outputWriter := bufio.NewWriter(output)
	for {
		content, _, err := inputReader.ReadLine()
		fmt.Printf("content:%s", content)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
		}
		outputString := string(content[2:5]) + "\r\n"
		fmt.Printf("   outputString:%s", outputString)
		_, writeErr := outputWriter.WriteString(outputString)
		outputWriter.Flush()
		if writeErr != nil {
			fmt.Println(writeErr)
			return
		}
	}
}
