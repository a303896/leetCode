package lib

import "fmt"

type Human struct {
	Name string
}

type Teacher struct {
	Name string
	Age int
	Human
}

type Student struct {
	Name string
	Age int
	Human
}

type A int

type B struct {
	num int
	Name string
}

//输出0-9
func Test() {
	LABEL:
		for i := 0; i < 10; i++ {
			for {
				fmt.Println(i)
				continue LABEL
			}
		}
}

//无限循环输出0
func Test2() {
	LABEL:
		for i := 0; i < 10; i++ {
			for {
				fmt.Println(i)
				goto LABEL
			}
		}
}

func Test3() {
	a := [...]int{1,2,3,4,5,6,7,8,9}
	sa := a[:]
	fmt.Println(sa)
}

func Test4() {
	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	fmt.Println(m)
	m2 := make(map[string]int)
	for k,v := range m {
		m2[v] = k
	}
	fmt.Println(m2)
}

func Test5() {
	fs := [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() {
			fmt.Println("defer func i=", i)
		}()
		fs[i] = func() {
			fmt.Println("func i=", i)
		}
	}
	for _,f := range fs {
		f()
	}
}

func Test6() {
	t := Teacher{Name:"王老师", Age:36, Human:Human{"中国人"}}
	s := Student{Name:"Lily", Age:18, Human:Human{"American"}}
	fmt.Println(t, s)
	fmt.Println(t.Human.Name)
}

func (a *A) Test7(num int) {
	*a+=A(num)
}

func (b *B) Test8() {
	b.num = 8
	fmt.Println(b)
}