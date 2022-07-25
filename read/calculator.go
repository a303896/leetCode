package read

import (
	"bufio"
	"fmt"
	"leetCode/lib"
	"os"
	"strconv"
)

/**
编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束
 */

func Calcu() {
	stack := new(lib.StructStack)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("input error")
			os.Exit(1)
		}
		input = input[0 : len(input)-1]
		switch {
			case input == "q":
				fmt.Println("Bye bye...")
				return
			case input >= "0" && input <= "99999":
				i,_ := strconv.Atoi(input)
				stack.Push(i)
				fmt.Printf("after push:%+v\n", stack)
			case input == "+":
				q := stack.Pop()
				p := stack.Pop()
				fmt.Printf("%d + %d = %d\n", p, q, p + q)
			case input == "-":
				q := stack.Pop()
				p := stack.Pop()
				fmt.Printf("%d - %d = %d\n", p, q, p - q)
			case input == "*":
				q := stack.Pop()
				p := stack.Pop()
				fmt.Printf("%d * %d = %d\n", p, q, p * q)
			case input == "/":
				q := stack.Pop()
				p := stack.Pop()
				fmt.Printf("%d / %d = %d\n", p, q, p / q)
			default:
				fmt.Println("invalid input")
		}
	}
}

/**
output:
&{Index:2 Data:[3 7 0 0]}
3 + 7 = 10
5
after push:&{Index:1 Data:[5 0 0 0]}
9
after push:&{Index:2 Data:[5 9 0 0]}
*
5 * 9 = 45
 */