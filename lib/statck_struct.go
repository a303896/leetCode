package lib

const LIMIT = 4

type StructStack struct {
	Index int
	Data [LIMIT]int
}

func (s *StructStack) Push(item int) {
	s.Data[s.Index] = item
	s.Index ++
}

func (s *StructStack) Pop() int {
	s.Index --
	tmp := s.Data[s.Index]
	s.Data[s.Index] = 0
	return tmp
}

//example
/*
s := new(lib.StructStack)
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
&{1 [10 0 0 0]}
&{2 [10 20 0 0]}
&{3 [10 20 30 0]}
a=30, s=&{2 [10 20 0 0]}
b=20, s=&{1 [10 0 0 0]}
&{2 [10 100 0 0]}
 */