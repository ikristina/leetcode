package problems

type NodeMinStack struct {
    min int
    val int
}

type MinStack struct {
    stack []NodeMinStack
}

func ConstructorMinStack() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int)  {
    newMin := val
    if len(this.stack) > 0 {
        curMin := this.stack[len(this.stack)-1].min
        if curMin < newMin {
            newMin = curMin
        }
    }

    this.stack = append(this.stack, NodeMinStack{val: val, min:newMin})
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0
    }
    return this.stack[len(this.stack)-1].val
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].min
}
