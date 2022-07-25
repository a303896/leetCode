package routine

func Sum(num []int, ch chan int) {
	sum := 0
	for _, i := range num {
		sum += i
	}
	ch <- sum
}


/**
	ch := make(chan int)
	go routine.Sum([]int{4,7,9,10}, ch)
	fmt.Println(<- ch)

output:
	30
 */