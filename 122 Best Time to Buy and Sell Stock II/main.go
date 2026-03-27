func maxProfit(prices []int) int {
    type State struct {
        idx int
        is_buy bool
    }

    var dp = make(map[State]int)
    var dfs func (i int, is_buy bool) int
    dfs = func (i int, is_buy bool) int {
        if i == len(prices) {
            return 0
        }
        state := State{idx: i, is_buy: is_buy}
        if _, exist := dp[state]; exist {
            return dp[state]
        }
        var sign int
        if is_buy {
            sign = -1
        } else {
            sign = 1
        }
        dp[state] = max(dfs(i+1, is_buy), dfs(i+1, !is_buy) + sign * prices[i])
        return dp[state]

    }
    return dfs(0, true)


    start := 0
    maxProfit := 0
    for end:=1; end < len(prices); end++ {
        if prices[end] > prices[start] {
            maxProfit += prices[end] - prices[start]
        }
        start = end

    }
    return maxProfit
}