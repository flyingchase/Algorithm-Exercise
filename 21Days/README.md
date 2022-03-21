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





## 2nd

25. K 个一组翻转链表【困难】https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

- 先写出在 l,r 区间内逆转链表的操作：
  - dummyHead 和 `i:=l;i<r;i++`
- 在整个链表中循环调用 `reverseBetween(head,l,r)`
  - 注意边界和循环条件：`for i:=1;i*k<=count;i++ { head=reverseBetween(head,(i--1)*k+1,i*k)}`

25. 三数之和【中等】https://leetcode-cn.com/problems/3sum/

- 注意非重复的优化

25. 排序数组（此题请至少尝试2种及以上的排序算法）【中等】https://leetcode-cn.com/problems/sort-an-array

- 1st 的 215 已完成常见三种 nlogn 算法的实现





## 3rd



1. 两数之和【简单】https://leetcode-cn.com/problems/two-sum/

   简单遍历即可

2. 买卖股票的最佳时机【简单】https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

3. 环形链表【简单】https://leetcode-cn.com/problems/linked-list-cycle/

   快慢指针

4. 剑指 Offer 26. 树的子结构【中等】https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

   递归，注意条件

   `a==nil||b==nil false;a.val==b.val&&helper(a.left,b.left)&&helper(a.right,b.right) true;returen isSub(a.l,b)||isSub(a.r,b)`

5. 二叉树的层序遍历【中等】https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

   宽度优先遍历即可















## 4th

剑指 Offer 14- I. 剪绳子【中等】https://leetcode-cn.com/problems/jian-sheng-zi-lcof/



160. 相交链表【简单】https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

     双指针，先得到两表长度，长链先走 gap，再同步走直至相遇，返回相遇结点/nil

161. 二叉树的锯齿形层序遍历【中等】https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

     宽度优先遍历，维护 flag 变量不断逆转 tmp

162. 有效的括号【简单】https://leetcode-cn.com/problems/valid-parentheses/

     栈：左括号入栈，右括号与栈顶匹配；最后返回栈空否

163. 合并两个有序数组【简单】https://leetcode-cn.com/problems/merge-sorted-array/

     merge 归并排序





## 5th

剑指 Offer 43. 1～n 整数中 1 出现的次数【困难】https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/

​	

236. 二叉树的最近公共祖先【中等】https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

     递归，注意判断：

     ``` go
     if l!=nil {
         if r!=nil {
             return root
         }
         return l
     }
     return r
     ```

     hashmap 存储父节点，从 pq 分别向上遍历，碰见相同结点则为公共祖先。

237. 搜索旋转排序数组【中等】https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

左侧有序在内查找







## 6th

5. 最长回文子串【中等】https://leetcode-cn.com/problems/longest-palindromic-substring/

   中心轴对称

   滑动窗口

   dp

6. 岛屿数量【中等】https://leetcode-cn.com/problems/number-of-islands/

   dfs 搜索 bfs

7. 全排列【中等】https://leetcode-cn.com/problems/permutations/

   回溯 注意去重

8. 字符串相加【简单】https://leetcode-cn.com/problems/add-strings/


​		注意余数 `carry=(n1+n2+carry)/10` 每次入结果时`res=strconv.Itoa(sum%10)+res`逆











## 7th







## 8th

199. 二叉树的右视图【中等】https://leetcode-cn.com/problems/binary-tree-right-side-view/

     宽度优先遍历，每次保留这一层 queue 的最后一个值即可

200. 重排链表【中等】https://leetcode-cn.com/problems/reorder-list/

     

201. 二叉树的中序遍历【简单】https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

     递归左根右，迭代 stack cur 到达最左结点过程中不断入栈

202. 用栈实现队列【简单】https://leetcode-cn.com/problems/implement-queue-using-stacks/

     双栈实现队列，in 栈和 out 栈

203. 二分查找【简单】https://leetcode-cn.com/problems/binary-search/

     板子题`l<=r`














