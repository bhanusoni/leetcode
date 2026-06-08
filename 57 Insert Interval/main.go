func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    for idx, interval := range intervals {
        if interval[1] < newInterval[0] {
            result = append(result, interval)
        } else if interval[0] > newInterval[1] {
            result = append(result, newInterval)
            return append(result, intervals[idx:]...)
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }
    return append(result, newInterval)
}