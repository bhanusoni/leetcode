type KthLargest struct {
    minHeap *priorityqueue.Queue
    k int
}


func Constructor(k int, nums []int) KthLargest {
    minHeap := priorityqueue.NewWith(utils.IntComparator)
    for _, num := range nums {
        minHeap.Enqueue(num)
        if minHeap.Size() > k {
            minHeap.Dequeue()
        }
    }
    return KthLargest{minHeap: minHeap, k: k}
}


func (this *KthLargest) Add(val int) int {
    this.minHeap.Enqueue(val)
    if this.minHeap.Size() > this.k {
        this.minHeap.Dequeue()
    }
    value, _ := this.minHeap.Peek()
    return value.(int)
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */