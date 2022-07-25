package read

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/**
读取文件内容  并在每一行头部加上行号
 */

func Cat(r *bufio.Reader) {
	i := 1
	for {
		content,err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%d             %s", i, content)
		if err == io.EOF {
			break
		}
		i++
	}
	return
}

func MainExec() {
	flag.Parse()
	if flag.NArg() == 0 {
		Cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(DIR + flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s 读取文件 %s 发生错误：%s", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		Cat(bufio.NewReader(f))
		f.Close()
	}
}

/**
go run .\main.go -n products.txt
1             "The ABC of Go";25.5;1500
2             "Functional Programming with Go";56;280
3             "Go for It";45.9;356
4             "The Go Way";55;500
 */
