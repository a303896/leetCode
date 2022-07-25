package lib

//两个有序数组合并成一个有序数组
func MergeSort(s1,s2 []int) []int {
	result := make([]int, 0, 0)
	l1, l2 := len(s1), len(s2)
	tmp := 0
	if l1 == 0 {
		return s2
	}
	if l2 == 0 {
		return s1
	}
	for j := 0; j < l1; j++ {
		for i := tmp; i < l2; i++ {
			if s1[j] < s2[i] {
				result = append(result, s1[j])
				tmp = i
				break
			}else {
				result = append(result, s2[i])
				if tmp == l2 - 1 {
					return append(result, s1[j : l1]...)
				}
			}
		}
	}
	if tmp < l2 - 1 {
		result = append(result, s2[tmp : l2]...)
	}
	return result
}