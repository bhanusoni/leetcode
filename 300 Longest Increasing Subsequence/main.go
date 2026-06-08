func lengthOfLIS(nums []int) int {
    LIS := []int{}
    for _, num := range nums {
        if len(LIS) == 0 || LIS[len(LIS)-1] < num {
            LIS = append(LIS, num)
        } else {
            left, right := 0, len(LIS)-1
            for left < right {
                mid := left + (right-left)/2
                if LIS[mid] < num {
                    left = mid + 1
                } else {
                    right = mid
                }
            }
            LIS[right] = num
        }
    }
    return len(LIS)
}