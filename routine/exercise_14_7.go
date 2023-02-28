package routine

func Exec14_7(ch,ch1 chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	ch1 <- 1
}


/**
example:
	ch := make(chan int)
	ch1 := make(chan int)
	go routine.Exec14_7(ch, ch1)
	for {
		select {
			case v := <-ch :
				fmt.Printf("ch 读取数字 %d \n", v)
				case <-ch1 :
					os.Exit(1)
		}
	}

output:
	ch 读取数字 0
	ch 读取数字 1
	ch 读取数字 2
	ch 读取数字 3
	ch 读取数字 4
	ch 读取数字 5
	ch 读取数字 6
	ch 读取数字 7
	ch 读取数字 8
	ch 读取数字 9
	ch 读取数字 10
	ch 读取数字 11
	ch 读取数字 12
	ch 读取数字 13
	ch 读取数字 14
*/