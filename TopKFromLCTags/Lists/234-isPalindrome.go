package topkLists

// stack to judge 注意奇数偶数
// 先遍历确定长度，再遍历一半入栈，继续后移并不断弹出栈顶比较
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	cur := head
	count := 1
	// 判断 cur.Next!=nil 确保 count 计数为链表长且 cur 指向最后一个结点
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	isOdd := false
	if count%2 != 0 {
		isOdd = true
	}
	stack := make([]int, count/2+1)
	cur = head
	count /= 2
	for count > 0 {
		count--
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	if isOdd {
		cur = cur.Next
	}
	for cur != nil {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != cur.Val {
			return false
		}
		cur = cur.Next
	}
	return true
}

// 双指针一遍遍历，slow 指针一边遍历一遍将左半部分逆序
// 注意判断 fast!=nil 则为奇数，slow 越过中间结点
func isPalindrome2(head *ListNode) bool {
	prev := &ListNode{Next: head}
	slow, fast := head, head
	// 注意判断 fast
	for fast != nil && fast.Next != nil {
		slow.Next, prev, slow, fast = prev, slow, slow.Next, fast.Next.Next
	}
	// fast 走两步最后不为空则为奇数
	// 奇数则 slow 跳过一步
	if fast != nil {
		slow = slow.Next
	}
	// slow 前半部分已逆序，逐个比较即可
	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true
}
