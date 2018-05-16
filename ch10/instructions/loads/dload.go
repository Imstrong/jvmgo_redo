package loads

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)
func _dload(frame *runtime.Frame,index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
//加载指令从局部变量表获取变量，推入操作数栈顶
type DLOAD struct {
	base.Index8Instruction
}
func (self *DLOAD) Execute(frame *runtime.Frame) {
	_dload(frame,self.Index)
}
type DLOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *DLOAD_0) Execute(frame *runtime.Frame) {
	_dload(frame,0)
}
type DLOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *DLOAD_1) Execute(frame *runtime.Frame) {
	_dload(frame,1)
}
type DLOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *DLOAD_2) Execute(frame *runtime.Frame) {
	_dload(frame,2)
}
type DLOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *DLOAD_3) Execute(frame *runtime.Frame) {
	_dload(frame,3)
}
