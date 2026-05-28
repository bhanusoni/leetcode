type Stock struct {
    price int
    isBuy bool
}

func maxProfit(prices []int) int {
    var dp = make(map[Stock]int)
    var dfs func(i int, isBuy bool) int
    dfs = func(i int, isBuy bool) int {
        if i >= len(prices) {
            return 0
        }
        stock := Stock{i, isBuy}
        if price, exist := dp[stock]; exist {
            return price
        }
        if isBuy {
            dp[stock] = max(
                dfs(i+1, !isBuy) - prices[i],
                dfs(i+1, isBuy),
            )
        } else {
            dp[stock] = max(
                dfs(i+2, !isBuy) + prices[i],
                dfs(i+1, isBuy),
            )
        }
        return dp[stock]
    }
    return dfs(0, true)
}