package routine

func Generate(ch chan int) {
	for i := 2; i < 1000 ; i++ {
		ch <- i
	}
	close(ch)
}

func Filter(in,out chan int, prime int) {
	for {
		i := <-in
		if i % prime != 0{
			out <- i
		}
	}
}


/**
	ch := make(chan int)
	go routine.Generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, "  ")
		ch1 := make(chan int)
		go routine.Filter(ch, ch1, prime)
		ch = ch1
	}
output:
	2  3  5  7  11  13  17  19  23  29  31  37  41  43  47  53  59  61  67  71  73  79  83  89  97  101  103  107...	//1000以内的所有素数
 */