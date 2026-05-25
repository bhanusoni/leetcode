class Solution:
    def myPow(self, x: float, n: int) -> float:
        def recursion(base, power):
            if power == 0:
                return 1
            if power%2 == 1:
                return base * recursion(base*base, (power-1)//2)
            return recursion(base*base, power//2)
        value = recursion(x, abs(n))
        return value if n >= 0 else 1/value
