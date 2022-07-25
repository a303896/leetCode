package lib

func Rotate(nums []int, k int)  {
	nums[k] = 10
	//if k == 0 {
	//	return
	//}
	//l := len(nums)
	//if k > l {
	//	k = k % l
	//}
	//nums = append(nums[l-k:], nums[0:l-k]...)
	//newNums := append(nums[l-k:], nums[0:l-k]...)
	//copy(nums,newNums)
}

var row = []int{-1,0,1,0}
var col = []int{0,-1,0,1}
func updateMatrix(mat [][]int) [][]int {
	w,h,result := len(mat),len(mat[0]),[][]int{}
	for index,value := range mat {
		for i,v := range value {
			if v == 0 {
				result[index][i] = 0
			}else {
				tmp := dfs(index,i,w,h,mat)
				result[index][i] = tmp[0] - index + tmp[1] - i
			}
		}
	}
	return result
}

func dfs(r,c,w,h int, mat [][]int) []int{
	result := []int{}
	for i := 0; i < 4; i++ {
		x,y := r + row[i], c + col[i]
		if x >= 0 && y >= 0 && x < w && y < h {
			if mat[x][y] == 0 {
				return []int{x,y}
			}
			result = dfs(x,y,w,h,mat)
		}
	}
	return result
}
