package loads

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)
func _iload(frame *runtime.Frame,index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
//加载指令从局部变量表获取变量，推入操作数栈顶
type ILOAD struct {
	base.Index8Instruction
}
func (self *ILOAD) Execute(frame *runtime.Frame) {
	_iload(frame,self.Index)
}
type ILOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_0) Execute(frame *runtime.Frame) {
	_iload(frame,0)
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_1) Execute(frame *runtime.Frame) {
	_iload(frame,1)
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_2) Execute(frame *runtime.Frame) {
	_iload(frame,2)
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *ILOAD_3) Execute(frame *runtime.Frame) {
	_iload(frame,3)
}
