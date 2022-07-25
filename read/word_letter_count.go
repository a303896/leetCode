package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数
 */

var (
	chars, words, lines = 0, 0, 0
)

func ReadInput() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if input == "S\n" {
			fmt.Printf("字符数：%d，字数：%d，行数：%d", chars, words, lines)
			os.Exit(0)
		}
		counters(input)
	}
}

func counters(str string)  {
	chars += len(str) - 1
	words += len(strings.Fields(str))
	lines ++
}

/**
fdsafds
asfsafa
ff
adf
S
字符数：19，字数：4，行数：4
 */