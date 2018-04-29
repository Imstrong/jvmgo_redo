package loads

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)
func _fload(frame *runtime.Frame,index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
//加载指令从局部变量表获取变量，推入操作数栈顶
type FLOAD struct {
	base.Index8Instruction
}
func (self *FLOAD) Execute(frame *runtime.Frame) {
	_fload(frame,self.Index)
}
type FLOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *FLOAD_0) Execute(frame *runtime.Frame) {
	_fload(frame,0)
}
type FLOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *FLOAD_1) Execute(frame *runtime.Frame) {
	_fload(frame,1)
}
type FLOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *FLOAD_2) Execute(frame *runtime.Frame) {
	_fload(frame,2)
}
type FLOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *FLOAD_3) Execute(frame *runtime.Frame) {
	_fload(frame,3)
}