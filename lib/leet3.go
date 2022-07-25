package lib

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)
	l := len(s)
	for index,_ := range s {
		for i := index + 1; i < l; i++ {
			if s[index] > s[i] {
				fmt.Printf("s[index]:%d--s[i]:%d\n", s[index], s[i])
				s[index],s[i] = s[i],s[index]
			}
			fmt.Printf("%v\n", s)
		}
	}
	mid := l / 2
	if l % 2 == 0 {
		return float64(s[mid - 1] + s[mid]) / 2
	}else {
		return float64(s[mid])
	}
}
