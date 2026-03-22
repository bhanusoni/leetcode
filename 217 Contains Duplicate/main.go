func containsDuplicate(nums []int) bool {
    unique := map[int]bool{}
    for _, n := range nums {
        if unique[n] {
            return true
        }
        unique[n] = true
    }
    return false
}