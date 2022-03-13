# 21Days Alg Training

## 1st

### 注意点

53.最大子数组和【简单】https://leetcode-cn.com/problems/maximum-subarray/

- Dp[i]代表 0-i中最大子数组之和，res:=nums[0]  dp[0]=nums[0]初始化 res 和开始



206.反转链表【简单】https://leetcode-cn.com/problems/reverse-linked-list/

- `var prev *ListNode = nil` 否则会多出一位 prev 默认0



21.合并两个有序联表【简单】https://leetcode-cn.com/problems/merge-two-sorted-lists/

- merge 注意 `copy(helper[index:],nums[pl:mid+1])  copy(helper[i:],n[pr:r+1]) copy(n[l:r+1],helper[:])`



3.无重复字符的最长子串【中等】https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

- Dict:=map[byte]int l,r 滑动 res=r-l 不用-1



215.数组中的第K个最大元素【中等】https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

