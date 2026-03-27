func moveZeroes(nums []int)  {
    i := 0
    for idx, num := range nums {
        if num == 0 {
            continue
        }
        if i != idx {
            nums[i], nums[idx] = nums[idx], nums[i]
        }
        i += 1
    }
}