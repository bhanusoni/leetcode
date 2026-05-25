class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()

        for i in range(len(nums)):
            if i != 0 and nums[i] == nums[i-1]:
                continue

            start = nums[i]
            left, right = i+1, len(nums)-1
            while left < right:
                value = start + nums[left] + nums[right]
                if value > 0:
                    right -= 1
                elif value < 0:
                    left += 1
                else:
                    res.append([start, nums[left], nums[right]])
                    left += 1
                    while left < right and nums[left] == nums[left-1]:
                        left += 1
        return res
                