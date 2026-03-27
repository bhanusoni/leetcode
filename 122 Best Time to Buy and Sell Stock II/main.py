class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        start = 0
        max_price = 0
        for end in range(1, len(prices)):
            if prices[start] < prices[end]:
                max_price += prices[end] - prices[start]
                start = end
            else:
                start = end
        return max_price
        dp = {}

        def dfs(i, is_buy):
            if i == len(prices):
                return 0
            if (i, is_buy) in dp:
                return dp[(i, is_buy)]
            sign = -1 if is_buy else 1
            dp[(i, is_buy)] = max(
                dfs(i + 1, not is_buy) + sign * prices[i],
                dfs(i + 1, is_buy)
            )
            return dp[(i, is_buy)]

        return dfs(0, True)
