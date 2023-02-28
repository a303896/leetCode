package routine

import "fmt"

func Pump1(ch chan int) {
	for i := 0; i < 100 ; i++ {
		ch <- i *2
	}
}

func Pump2(ch chan int) {
	for i := 0; i < 100 ; i++ {
		ch <- i + 5
	}
}

func SuckNew(ch1,ch2 chan int) {
	for {
		select {
		case v := <-ch1 :
			fmt.Printf("ch1获得数字：%d\n", v)
		case v := <-ch2 :
			fmt.Printf("ch2获得数字：%d\n", v)
		}
	}
}

/*
	ch1 := make(chan int)
	ch2 := make(chan int)
	go routine.Pump1(ch1)
	go routine.Pump2(ch2)
	go routine.SuckNew(ch1, ch2)

	time.Sleep(1e9)
 */

/*
output:
ch1获得数字：0
ch1获得数字：2
ch2获得数字：5
ch2获得数字：6
ch2获得数字：7
......
ch1获得数字：192
ch1获得数字：194
ch1获得数字：196
ch1获得数字：198
 */
