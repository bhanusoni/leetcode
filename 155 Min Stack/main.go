type MinStack struct {
    stk [][2]int
}


func Constructor() MinStack {
    return MinStack{
        stk: [][2]int{},
    }
}


func (this *MinStack) Push(val int)  {
    if len(this.stk) == 0 {
        this.stk = append(this.stk, [2]int{val, val})
    } else {
        this.stk = append(this.stk, [2]int{val, min(val, this.GetMin())})
    }
}


func (this *MinStack) Pop()  {
    this.stk = this.stk[:len(this.stk)-1]
}


func (this *MinStack) Top() int {
    return this.stk[len(this.stk)-1][0]
}


func (this *MinStack) GetMin() int {
    return this.stk[len(this.stk)-1][1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */