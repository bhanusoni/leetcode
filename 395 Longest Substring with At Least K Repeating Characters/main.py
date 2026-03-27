from collections import Counter


class Solution:
    def longestSubstring(self, s: str, k: int) -> int:
        if not s or k > len(s):
            return 0
        if k == 0:
            return len(s)

        l = 0
        count = Counter(s)
        while l < len(s) and count[s[l]] >= k:
            l += 1
        if l >= len(s):
            return len(s)

        ls1 = self.longestSubstring(s[:l], k)
        while l < len(s) and count[s[l]] < k:
            l += 1
        ls2 = self.longestSubstring(s[l:], k) if l < len(s) - 1 else 0

        return max(ls1, ls2)
