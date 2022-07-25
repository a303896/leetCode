package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type product struct {
	Title string
	Price float32
	Count int
}

func ReadCsv() {
	file := "./read/example/products.txt"
	input,err := os.Open(file)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	defer input.Close()
	reader := bufio.NewReader(input)
	products := make([]product, 0)
	for {
		content,error := reader.ReadString('\n')
		slice := strings.Split(content, ";")
		pro := product{}
		pro.Title = slice[0]
		price,_ := strconv.ParseFloat(slice[1], 32)
		pro.Price = float32(price)
		count,err1 := strconv.Atoi(strings.Replace(slice[2], "\r\n", "", -1))
		if err1 != nil {
			fmt.Printf("%v", err1)
		}else {
			//fmt.Printf("count=%d", count)
		}
		pro.Count = count
		products = append(products, pro)
		if error == io.EOF {
			fmt.Printf("%+v", products)
			return
		}
	}
}

/**
output:
[{Title:"The ABC of Go" Price:25.5 Count:1500} {Title:"Functional Programming with Go" Price:56 Count:280} {Title:"Go for It" Price:45.9 Count:356} {Title:"The Go Way" Price:55 Count:500}]
 */