func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }
    size := len(nums)
    i := 0
    var permute_values = [][]int{}
    for i < size {
        n := nums[0]
        nums = nums[1:]
        values := permute(nums)
        nums = append(nums, n)
        for idx, value := range(values) {
            values[idx] = append(value, n)
        }
        permute_values = append(permute_values, values...)
        i += 1
    }
    return permute_values
}