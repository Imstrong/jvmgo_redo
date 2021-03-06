package runtime

import "jvmgo/ch10/runtime/heap"

type Thread struct {
	//pc寄存器，保存当前字节码指令行号
	pc int
	//指向虚拟机栈的引用
	stack *Stack
}
func NewThread() *Thread {
	//创建Thread实例并指定最大帧数
	return &Thread{stack:newStack(1024)}
}
func(self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func(self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self,method)
}
func (self *Thread) SetPC(pc int) {
	self.pc=pc
}
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}
func (self *Thread) TopFrame() *Frame {
	return self.stack._top
}
func (self *Thread) ClearStack() {
	self.stack.clear()
}
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}
func (self *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}