package routine

import (
	"fmt"
)

func SendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func GetData(ch chan string) {
	var input string
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}