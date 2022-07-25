package lib

type Stack [4]int

func (s *Stack) Push(item int) {
	for i,val := range s {
		if val == 0 {
			s[i] = item
			break
		}
	}
}

func (s *Stack) Pop() int {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if s[i] != 0 {
			temp := s[i]
			s[i] = 0
			return temp
		}
	}
	return 0
}

//example
/*
s := new(lib.Stack)
s.Push(10)
fmt.Println(s)
s.Push(20)
fmt.Println(s)
s.Push(30)
fmt.Println(s)
a := s.Pop()
fmt.Printf("a=%d, s=%v\n", a, s)
b := s.Pop()
fmt.Printf("b=%d, s=%v\n", b, s)
s.Push(100)
fmt.Println(s)
*/

//output
/*
&[10 0 0 0]
&[10 20 0 0]
&[10 20 30 0]
a=30, s=&[10 20 0 0]
b=20, s=&[10 0 0 0]
&[10 100 0 0]
*/