package routine

func Integers() chan int {
	resume := make(chan int)
	count := 0
	go func() {
		for{
			resume <- count
			count ++
		}
	}()
	return resume
}

func GenerateInteger(ch chan int) int {
	return <- ch
}
