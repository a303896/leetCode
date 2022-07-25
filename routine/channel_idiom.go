package routine

import "fmt"

func Pump() chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0 ; ; i++ {
			ch <- i
		}
	}(ch)
	//go func() {
	//	for i := 0 ; ; i++ {
	//		ch <- i
	//	}
	//}()
	return ch
}

func Suck(ch chan int) {
	for {
		fmt.Println(<- ch)
	}
}

func Suck2(ch chan int) {
	go func() {
		for num := range ch {
			fmt.Println(num)
		}
	}()
}

/**
	ch := routine.Pump()
	go routine.Suck(ch)
	time.Sleep(1e9)
output:

 */