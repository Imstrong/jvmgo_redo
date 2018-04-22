package runtime
type Thread struct {
	//pc寄存器，保存当前字节码指令行号
	pc int
	//指向虚拟机栈的引用
	stack *Stack
}
func newThread() *Thread {
	//创建Thread实例并指定最大帧数
	return &Thread{stack:newStack(1024)}
}
func(self *Thread) pushFrame(frame *Frame) {
	self.stack.push(frame)
}
func(self *Thread) popFrame() {
	self.stack.pop()
}
func (self *Thread) currentFrame() *Frame{
	return self.stack.top()
}

