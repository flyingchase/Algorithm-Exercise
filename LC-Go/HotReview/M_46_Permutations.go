package HotReview
/*
无重复数字的全排列
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	used, lists, res := make([]bool, len(nums)), []int{}, [][]int{}
	dfs(nums,0, lists,&res,&used)
	return res
}

func dfs(nums []int, index int, lists []int, res *[][]int, used  *[]bool) {
	if index == len(nums) {
		// 新建 tmp 防止指针改变回溯的操作
		tmp:=make([]int,len(lists))
		copy(tmp, lists)
		*res= append(*res,tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i]=true
			lists =append(lists,nums[i])
			dfs(nums,index+1, lists,res,used)
			// 退回操作
			lists = lists[:len(lists)-1]
			(*used)[i]=false
		}
	}
}
