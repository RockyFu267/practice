package sample

//买卖股票的最佳时机
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

//maxProfit 时间94%，内存100%
func maxProfit(prices []int) int {
	lenNum := len(prices)
	var res int
	for k := 1; k < lenNum; k++ {
		if (prices[k-1] - prices[k]) < 0 {
			res = res + (prices[k-1] - prices[k])
		}
	}
	return res * -1
}
