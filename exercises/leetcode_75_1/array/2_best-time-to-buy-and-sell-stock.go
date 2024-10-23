/*
	Problem: 121. Best Time to Buy and Sell Stock
	Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
	Complexity: O(n)
*/

func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := -1
    for _, p := range prices {
        if minPrice > p || minPrice == -1 {
            minPrice = p
        }
        if p - minPrice > maxProfit {
            maxProfit = p - minPrice
        }
    }

    return maxProfit
}
