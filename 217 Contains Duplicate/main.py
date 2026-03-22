class Solution:
    def containsDuplicate(self, nums: list[int]) -> bool:
        unique = set(nums)
        return len(unique) != len(nums)
