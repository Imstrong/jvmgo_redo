package stores

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)

//存储指令从操作数栈弹出变量，并存入局部变量表
func _fstore(frame *runtime.Frame,index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index,val)
}
type FSTORE struct {base.Index8Instruction}
func (self *FSTORE) Execute(frame *runtime.Frame) {
	_fstore(frame,self.Index)
}
type FSTORE_0 struct {base.NoOperandsInstruction}
func (self *FSTORE_0) Execute(frame *runtime.Frame) {
	_fstore(frame,0)
}
type FSTORE_1 struct {base.NoOperandsInstruction}
func (self *FSTORE_1) Execute(frame *runtime.Frame) {
	_fstore(frame,1)
}
type FSTORE_2 struct {base.NoOperandsInstruction}
func (self *FSTORE_2) Execute(frame *runtime.Frame) {
	_fstore(frame,2)
}
type FSTORE_3 struct {base.NoOperandsInstruction}
func (self *FSTORE_3) Execute(frame *runtime.Frame) {
	_fstore(frame,3)
}
