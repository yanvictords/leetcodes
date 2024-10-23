/*
	Problem: 322. Coin Change
	Reference: https://leetcode.com/problems/coin-change/
	Complexity: O(N * Amount)
*/

func dfs(coins []int, amount int, answers map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt32
	}
	if calculatedMin, ok := answers[amount]; ok {
		return calculatedMin
	}

	answers[amount] = math.MaxInt32
	for _, coin := range coins {
		if amount >= coin {
			answers[amount] = min(
				answers[amount],
				1+dfs(coins, amount-coin, answers),
			)
		}
	}

	return answers[amount]
}

func coinChange(coins []int, amount int) int {
	answers := make(map[int]int, amount+1)
	dfs(coins, amount, answers)
	if answers[amount] == math.MaxInt32 {
		return -1
	}
	return answers[amount]
}