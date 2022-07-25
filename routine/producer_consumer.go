package routine

import "fmt"

func Produce(ch chan int) {
	for i := 0; i < 100; i += 10 {
		ch <- i
	}
	close(ch)
}

func Consume(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Println(num)
	}
	done <- true
}

/**
	ch := make(chan int)
	done := make(chan bool)
	go routine.Produce(ch)
	go routine.Consume(ch, done)
	<- done

output:
	0
	10
	20
	30
	40
	50
	60
	70
	80
	90
 */
