package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
程序中的数据结构如下，是一个包含以下字段的结构:

type Page struct {
    Title string
    Body  []byte
}
请给这个结构编写一个 save() 方法，将 Title 作为文件名、Body 作为文件内容，写入到文本文件中。

再编写一个 load() 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。请使用 *Page 做为参数，因为这个结构可能相当巨大，我们不想在内存中拷贝它
 */

const DIR =  "./read/example/"

type Page struct {
	Title string
	Body []byte
}

func (page *Page) Save() {
	file,error := os.OpenFile(DIR + page.Title, os.O_CREATE | os.O_APPEND, 0666)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()
	output := bufio.NewWriter(file)
	output.Write(page.Body)
	output.Flush()
}

func Load(page *Page) {
	file,error := os.Open(DIR + page.Title)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()
	input := bufio.NewReader(file)
	for{
		bt := make([]byte, 4)
		n, err := input.Read(bt)
		page.Body = append(page.Body, bt...)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(bt[:n]))
	}
}

/**

	page1 := read.Page{"test", []byte{}}
	page2 := read.Page{"test", []byte{'a', 'p', 'p', 'l', 'e'}}
	page3 := read.Page{"test", []byte{'I', 'B', 'M', 'A'}}
	page2.Save()
	page3.Save()
	read.Load(&page1)
	fmt.Println(page1.Body)

output:
4 appl
4 eIBM
4 Aapp
4 leIB
2 MA
[97 112 112 108 101 73 66 77 65 97 112 112 108 101 73 66 77 65 0 0 0 0 0 0]
*/