package lib

import "strings"

//https://leetcode-cn.com/problems/zigzag-conversion/
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var z [][]rune
	var rows int
	l := len(s)
	if numRows < l {
		rows = numRows
	}else {
		rows = l
	}
	for i := 0; i < rows; i++ {
		z = append(z, []rune{})
	}
	col := 0
	isBool := false
	for _,v := range s {
		z[col] = append(z[col], v)
		if col == 0 || col == rows - 1 {
			isBool = !isBool
		}
		if isBool {
			col++
		}else {
			col--
		}
	}
	var result string
	for _,item := range z {
		result += string(item)
	}
	return result
}

//https://leetcode-cn.com/problems/zigzag-conversion/solution/ji-jian-jie-fa-by-ijzqardmbd/
func Convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	n := numRows * 2 - 2
	for i,v := range s {
		m := i % n
		result[min(m, n - m)] += string(v)
	}
	return strings.Join(result, "")
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}