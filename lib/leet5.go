package lib

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/reverse-integer/
func Reverse(x int) int {
	fmt.Printf("%d\n", math.MinInt32)
	fmt.Printf("%d\n", math.MaxInt32)
	a := false
	if x < 0 {
		x = -x
		a = true
	}
	sx := strings.Split(strconv.Itoa(x), "")
	l := len(sx)
	for i := 0; i < l / 2; i++ {
		n := l - 1 - i
		sx[i], sx[n] = sx[n], sx[i]
	}
	res,_ := strconv.Atoi(strings.Join(sx, ""))
	if a {
		res = -res
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	fmt.Printf("%d", res)
	return res
}

func Reverse2(x int) int {
	var rev int
	for x != 0 {
		fmt.Printf("%d\n", rev)
		fmt.Printf("%d\n", x)
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		num := x % 10
		x /= 10
		rev = rev * 10 + num
	}
	return rev
}
