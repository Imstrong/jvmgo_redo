package loads

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch05/runtime"
)

type LLOAD struct {
	base.Index16Instruction
}
func _lload(frame *runtime.Frame,index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
func (self *LLOAD) Execute(frame *runtime.Frame) {
	_lload(frame,self.Index)
}
type LLOAD_0 struct {
	base.NoOperandsInstruction
}
func (self *LLOAD_0) Execute(frame *runtime.Frame) {
	_lload(frame,0)
}
type LLOAD_1 struct {
	base.NoOperandsInstruction
}
func (self *LLOAD_1) Execute(frame *runtime.Frame) {
	_lload(frame,1)
}
type LLOAD_2 struct {
	base.NoOperandsInstruction
}
func (self *LLOAD_2) Execute(frame *runtime.Frame) {
	_lload(frame,2)
}
type LLOAD_3 struct {
	base.NoOperandsInstruction
}
func (self *LLOAD_3) Execute(frame *runtime.Frame) {
	_lload(frame,3)
}
