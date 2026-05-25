class Solution:
    def nextPermutation(self, nums: list[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        pvt_idx = -1
        for i in range(len(nums)-2, -1, -1):
            if nums[i] < nums[i+1]:
                pvt_idx = i
                break
        if pvt_idx == -1:
            nums.sort()
            return
        for i in range(len(nums)-1, -1, -1):
            if nums[pvt_idx] < nums[i]:
                nums[pvt_idx], nums[i] = nums[i], nums[pvt_idx]
                break
        nums[pvt_idx+1:] = sorted(nums[pvt_idx+1:])
