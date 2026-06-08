class Solution:
    def minTaps(self, n: int, ranges: list[int]) -> int:
        nums = [0] * (n+1)
        for i, r in enumerate(ranges):
            low, high = max(0, i-r), i+r
            nums[low] = max(nums[low], high)
        
        taps = 0
        covered = 0
        next_covered = 0
        for i, num in enumerate(nums):
            if i > covered:
                return -1
            next_covered = max(next_covered, num)
            if i == covered:
                covered = next_covered
                taps += 1
            if covered >= n:
                return taps
        return taps if covered >= n else -1


        actual_ranges = [[max(0, idx-r), idx+r] for idx, r in enumerate(ranges) if r > 0]
        actual_ranges.sort(key=lambda x: (x[0], -x[1]))
        print(actual_ranges)
        # if not actual_ranges or actual_ranges[0][0] != 0:
        #     return -1
        covered_till, total = 0, 0
        next_covered_till = 0
        for i, (x, y) in enumerate(actual_ranges):
            if i > covered_till:
                return -1
            next_covered_till = max(next_covered_till, y)
            if i == covered_till:
                covered_till = next_covered_till
                total +=1
            if covered_till >= n:
                return total
        return -1
