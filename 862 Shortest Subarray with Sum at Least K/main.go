func shortestSubarray(nums []int, k int) int {
    res := len(nums) + 1
    elements := list.New()
    currSum := 0
    elements.PushBack([]int{0, 0})
    for idx, num := range nums {
        currSum += num
        for elements.Len() > 0 && elements.Back().Value.([]int)[1] > currSum {
            elements.Remove(elements.Back())
        }
        elements.PushBack([]int{idx+1, currSum})
        // fmt.Println(elements.Back().Value.([]int), elements.Front().Value.([]int), currSum)
        for currSum - elements.Front().Value.([]int)[1] >= k {
            front := elements.Front()
            res = min(res, idx - front.Value.([]int)[0]+1)
            elements.Remove(front)
        }
    }
    if res == len(nums) + 1 {
        return -1
    }
    return res
}