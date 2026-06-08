func lastStoneWeight(stones []int) int {
    var maxHeap = priorityqueue.NewWith(func (a, b interface{}) int {
        return b.(int) - a.(int)
    })

    for _, s := range stones {
        maxHeap.Enqueue(s)
    }

    for maxHeap.Size() > 1 {
        first, _ := maxHeap.Dequeue()
        second, _ := maxHeap.Dequeue()
        element := first.(int) - second.(int)
        if element != 0 {
            maxHeap.Enqueue(element)
        }
    }
    maxHeap.Enqueue(0)
    res,  _ :=  maxHeap.Dequeue()
    return res.(int)

}