package lib

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
}

type Person struct {
	Name string
	Age int
}

func (s *ListNode) String() string {
	return fmt.Sprintf("{Val: %v, Next: %+v}", s.Val, *s.Next)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var tmp *ListNode
	carry := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		a,b := 0,0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		if result == nil {
			result = &ListNode{Val: sum}
			tmp = result
		}else {
			tmp.Next = &ListNode{Val: sum}
			tmp = tmp.Next
		}
		if carry>0 {
			tmp.Next = &ListNode{Val: carry}
		}
		fmt.Printf("result：%+v\n", *result)
		fmt.Printf("tmp：%+v\n", *tmp)
	}
	fmt.Printf("%+v", *result)
	return result
}
