package control

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)
//没有返回值
type RETURN struct {base.NoOperandsInstruction}
//返回引用类型
type ARETURN struct {
	base.NoOperandsInstruction
}
//返回int类型
type IRETURN struct {
	base.NoOperandsInstruction
}
//返回long类型
type LRETURN struct {
	base.NoOperandsInstruction
}
//返回float类型
type FRETURN struct {
	base.NoOperandsInstruction
}
//返回double类型
type DRETURN struct{
	base.NoOperandsInstruction
}
func (self *RETURN) Execute(frame *runtime.Frame) {
	frame.Thread().PopFrame()
}
func (self *IRETURN) Execute(frame *runtime.Frame) {
	//获取当前帧的线程
	thread:=frame.Thread()
	//栈顶元素出栈，标识为当前帧
	currentFrame:=thread.PopFrame()
	//出栈后的栈顶元素为该方法调用者（上层方法）所有帧
	invokerFrame:=thread.TopFrame()
	retVal:=currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}
func (self *LRETURN) Execute(frame *runtime.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	returnValue:=currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(returnValue)
}
func (self *FRETURN) Execute(frame *runtime.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	returnValue:=currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(returnValue)
}
func (self *DRETURN) Execute(frame *runtime.Frame) {
	thread:=frame.Thread()
	currentFrame:=thread.PopFrame()
	invokerFrame:=thread.TopFrame()
	returnValue:=currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(returnValue)
}