package HotReview

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

买卖股票的最佳时机
*/
// 模拟 dp
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, lowest := 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[lowest] {
			lowest = i
		} else if cur := prices[i] - prices[lowest]; cur > max {
			max = cur
		}
	}
	return max
}



// 单调栈
func maxProfit1(prices []int) int {
	if len(prices)==0 {
		return 0
	}
	stack,res:=[]int{prices[0]},0
	for i:=1;i<len(prices); i++{
		if prices[i]>stack[len(stack)-1] {
			stack=append(stack,prices[i])
		}else {
			// 找到栈中的比当前 i小的最大值
			index:=len(stack) - 1
			for ; index>=0;index-- {
				if stack[index]<prices[i] {
					break
				}
			}
			// 丢弃所找到的比当前 i 小的最大值更大的部分
			stack=stack[:index+1]
			stack=append(stack,prices[i])
		}
		res= maxMaxProfix(res,stack[len(stack)-1]-stack[0])
	}
	return res
}

func maxMaxProfix(j int, i int) int {
	if i > j {
		return i
	}
	return j
}

















































































