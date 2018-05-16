package stores

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

//存储指令从操作数栈弹出变量，并存入局部变量表
func _istore(frame *runtime.Frame,index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index,val)
}
type ISTORE struct {base.Index8Instruction}
func (self *ISTORE) Execute(frame *runtime.Frame) {
	_istore(frame,self.Index)
}
type ISTORE_0 struct {base.NoOperandsInstruction}
func (self *ISTORE_0) Execute(frame *runtime.Frame) {
	_istore(frame,0)
}
type ISTORE_1 struct {base.NoOperandsInstruction}
func (self *ISTORE_1) Execute(frame *runtime.Frame) {
	_istore(frame,1)
}
type ISTORE_2 struct {base.NoOperandsInstruction}
func (self *ISTORE_2) Execute(frame *runtime.Frame) {
	_istore(frame,2)
}
type ISTORE_3 struct {base.NoOperandsInstruction}
func (self *ISTORE_3) Execute(frame *runtime.Frame) {
	_istore(frame,3)
}
