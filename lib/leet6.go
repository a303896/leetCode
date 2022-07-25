package lib

import (
	"math"
)

//https://leetcode-cn.com/problems/string-to-integer-atoi/
func MyAtoi(s string) int {
	result := 0
	neg := false
	each := true
	for _,ch := range []byte(s) {
		if each && ch == ' ' {
			continue
		}else if(each && ch == '-') {
			neg = true
			each = false
		}else if(each && ch == '+') {
			each = false
		}else {
			each = false
			ch -= '0'
			if ch <0 || ch > 9 {
				break
			}
			result = result * 10 + int(ch)
			if result > math.MaxInt32 {
				break
			}
		}
	}
	if neg {
		result = -result
	}
	if result < math.MinInt32 {
		result = math.MinInt32
	}
	if result > math.MaxInt32 {
		result = math.MaxInt32
	}
	return result
}
