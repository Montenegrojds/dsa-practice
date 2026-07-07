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

type Node struct{
	val int
	min int
	next *Node
}
type MinStack struct {
	top *Node
}

func Constructor() MinStack {
	return MinStack{top: nil}
}

func (this *MinStack) Push(val int) {
	if this.top == nil{
		newNode := &Node{val:val,min:val,next: this.top}
		this.top= newNode

	}else{
		if val<this.top.min{
			newNode := &Node{val: val,min: val,next: this.top}
			this.top=newNode
		}else{
			newNode := &Node{val: val, min:this.top.min ,next: this.top}
			this.top=newNode
			}	
	}
}

func (this *MinStack) Pop() {
	if this.top==nil{
		fmt.Println("Empty stack")
	}else{
	node:= this.top.next
	this.top = node
	}
}

func (this *MinStack) Top() int {
return this.top.val
}

func (this *MinStack) GetMin() int {
return this.top.min
}
