package routine

import "fmt"

func Fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	}else {
		v1, _ := Fibonacci(n - 1)
		v2, _ := Fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	fmt.Printf("Fibonacci---val=%d,pos=%d\n", val, pos)
	return
}

//func GoFibonacci(ch, ch1 chan int, n int) {
//
//	if n <= 1 {
//		ch <- 1
//	}else {
//		v1, _ := GoFibonacci(ch, ch1, n - 1)
//		v2, _ := GoFibonacci(ch, ch1, n - 2)
//		ch <- v1 + v2
//	}
//	ch1 <- 1
//}