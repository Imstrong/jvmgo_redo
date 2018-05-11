package loads

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
)
func _aload(frame *runtime.Frame,index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
//加载指令从局部变量表获取变量，推入操作数栈顶
type ALOAD struct {
	base.Index8Instruction
}
func (self *ALOAD) Execute(frame *runtime.Frame) {
	_aload(frame,self.Index)
}
type ALOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *ALOAD_0) Execute(frame *runtime.Frame) {
	_aload(frame,0)
}
type ALOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *ALOAD_1) Execute(frame *runtime.Frame) {
	_aload(frame,1)
}
type ALOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *ALOAD_2) Execute(frame *runtime.Frame) {
	_aload(frame,2)
}
type ALOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *ALOAD_3) Execute(frame *runtime.Frame) {
	_aload(frame,3)
}
