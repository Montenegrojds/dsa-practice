/*
Min Stack
Medium Topics Company Tags
Hints

Design a stack class that supports the push, pop, top, and getMin operations.

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

Each function should run in O(1)O(1) time.

Example 1:

Input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]

Output: [null,null,null,null,0,null,2,1]

Explanation:
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1
*/

type MinStack struct {
	stack    []int 
    minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}

}

func (this *MinStack) Push(val int) {
	this.stack=append(this.stack,val)
	if len(this.minStack)==0{
		this.minStack=append(this.minStack,val)
	}else{
		last:=len(this.minStack)-1
		if this.minStack[last]<val{
			this.minStack=append(this.minStack,this.minStack[last])
		}else{
			this.minStack=append(this.minStack,val)
		}

	}
	

}

func (this *MinStack) Pop() {
	last:=len(this.stack)-1
	if len(this.stack)==0{
		fmt.Println("Empty stack")
		return
	}else{
		this.stack=this.stack[:last]
		this.minStack=this.minStack[:last]
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.stack)-1]
}
