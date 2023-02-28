package main

import "leetCode/routine"

func main() {
	//l1 := lib.ListNode{2, &lib.ListNode{4, &lib.ListNode{3, nil}}}
	//l2 := lib.ListNode{5, &lib.ListNode{6, &lib.ListNode{4, nil}}}
	//lib.AddTwoNumbers(&l1, &l2)

	//p1 := lib.Person{"lilei", 12}
	//p2 := p1
	//p2.Name = "HanMei"
	//fmt.Printf("%+v\n", p1)
	//fmt.Printf("%+v\n", p2)
	//s1 := []int{1,2,3}
	//s2 := s1
	//s2[1] = 4
	//fmt.Printf("%v\n", s1)
	//fmt.Printf("%v", s2)

	//lib.FindMedianSortedArrays([]int{3}, []int{-2,-1})

	//lib.Convert("PAYPALISHIRING", 3)
	//println(lib.Reverse2(1534236469))
	//fmt.Printf("%d", lib.MyAtoi("9223372036854775808"))
	//nums := []int{1,2,3,4,5,6,7}
	//lib.Rotate(nums, 3)
	//fmt.Printf("%v",nums)
	//lib.Test()
	//lib.Test2()
	//lib.Test3()
	//lib.Test4()
	//lib.Test5()
	//lib.Test6()
	//a := lib.A(1)
	//fmt.Println("a=", a)
	//a.Test7(200)
	//fmt.Println("a=", a)
	//b := lib.B{
	//	Name: "bbb",
	//}
	//b.Test8()

	//反射示例测试
	//user := lib.User{"懂法守法", 18}
	//user.Info()
	//admin := lib.Admin{User:lib.User{"administrator", 18}, Title: "超级管理员"}
	//lib.GetAnonymousField(admin)
	//driver := lib.Driver{"土豪A", 18}
	//driver.Drive(lib.BMW{"宝马X5", "红色", 500000})
	//v := reflect.ValueOf(driver)
	//m := v.MethodByName("Drive")
	//m.Call([]reflect.Value{reflect.ValueOf(lib.BMW{"宝马X3", "黑色", 200000})})
	//lib.NewDlinkedList([]int{101,102,103})
	//lib.GetSize()
	//lib.PatternTest("5.20 is a stupid festival like 11.11", "\\d+\\.\\d+")
	//d := lib.Day(5)
	//fmt.Println(d)

	//read.Exec()
	//read.Execute()
	//read.ReadInput()
	//s1 := []int{1,3,5,6}
	//s2 := []int{2,4,7,8}
	//fmt.Println(lib.MergeSort(s1, s2))
	//s3 := []int{}
	//s4 := []int{2,5,8,9}
	//fmt.Println(lib.MergeSort(s3, s4))
	//s5 := []int{1,4,6,8,10}
	//s6 := []int{2,5}
	//fmt.Println(lib.MergeSort(s5, s6))
	//s7 := []int{3,7}
	//s8 := []int{1,4,8,12,15}
	//fmt.Println(lib.MergeSort(s7, s8))
	//read.Connect()
	//read.Calcu()
	//read.ReadCsv()
	//read.Hello()
	//read.Echo()
	//read.MainExec()
	//read.ExecRemove()
	//read.Encode()
	//read.Decode()

	//ch := make(chan string)
	//go routine.SendData(ch)
	//go routine.GetData(ch)
	//time.Sleep(1e9)

	//out := make(chan int, 100)
	//go func(ch chan int) {
	//	for i := 0; i < 100; i++{
	//		ch <- i
	//	}
	//}(out)
	//for i := 0; i < 100; i++{
	//	out <- i
	//}
	//go func(ch chan int) {
	//	for {
	//		fmt.Println(<-out)
	//	}
	//}(out)
	//time.Sleep(1e9)

	//ch := routine.Sieve()
	//fmt.Print(<-ch, " ")

	//res, pos := routine.Fibonacci(10)
	//fmt.Printf("res=%d,pos=%d", res, pos)
	//routine.TickTok()
	//routine.Start()
	//routine.Start2()
	//routine.Start3()
	routine.PrintOnce()
}
