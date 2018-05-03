package runtime

import "jvmgo/ch07/runtime/heap"

//jvm栈帧
type Frame struct {
	//下一帧的指针
	lower *Frame
	//局部变量表引用
	localVars LocalVars
	//	操作数栈指针
	operandStack *OperandStack
	thread *Thread
	method *heap.Method
	nextPC int
}
func NewFrame(thread *Thread,method *heap.Method) *Frame {
	return &Frame{
		//maxLocals定义了局部变量表大小maxStack定义操作数栈深度，由编译器计算并保存在方法表的Code属性中，在上一章已经能拿到了
		localVars:newLocalVars(method.MaxLocals()),
		operandStack:newOperandStack(method.MaxStack()),
		thread:thread,
	}
}
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(pc int) {
	self.nextPC=pc
}
func (self *Frame) Method() *heap.Method{
	return self.method
}
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method: method,
		localVars: newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
