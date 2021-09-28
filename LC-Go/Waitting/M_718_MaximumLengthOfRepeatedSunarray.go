package Waitting

/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
*/

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1)*len(nums2)==0 {
		return 0
	}
	l1:=len(nums1)
	l2:=len(nums2)
	dp:=make([][]int,l1+1)
	for i := range dp {
	    dp[i] = make([]int,l2+1)
	}

	for i:=1;i<=l1; i++ {
		for j:=1;j<=l2;l2++{
			if nums1[i-1]==nums2[j-1] {
				dp[i][j]=dp[i-1][j-1]+1
			}else {
			    dp[i][j]=max718(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func max718(i int, j int) int {
	if	i < j {
	    return j
	}
	return i
}


