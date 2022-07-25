package routine

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 1000; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	ch := make(chan int)
	go func() {
		//for i := range in {
		//	if i % prime != 0 {
		//		ch <- i
		//	}
		//}
		for {
			if i := <- in; i % prime != 0 {
				ch <- i
			}
		}
	}()
	return ch
}

func Sieve() chan int {
	ch := make(chan int)
	go func() {
		in := generate()
		for {
			prime := <-in
			in = filter(in, prime)
			ch <- prime
		}
	}()
	return ch
}