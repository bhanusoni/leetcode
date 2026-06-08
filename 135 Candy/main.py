class Solution:
    def candy(self, ratings: list[int]) -> int:
        n = len(ratings)
        candies = [1]*n
        res = 0
        for i in range(1, n):
            if ratings[i] > ratings[i-1]:
                candies[i] = candies[i-1] + 1
        for i in range(n-2, -1, -1):
            if ratings[i] > ratings[i+1]:
                candies[i] = max(1+candies[i+1], candies[i])
            res += candies[i]
        return res + candies[-1]