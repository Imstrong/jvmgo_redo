package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch04/runtime"
)

//存储指令从操作数栈弹出变量，并存入局部变量表
func _lstore(frame *runtime.Frame,index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index,val)
}
type LSTORE struct {base.Index8Instruction}
func (self *LSTORE) Execute(frame *runtime.Frame) {
	_lstore(frame,self.Index)
}
type LSTORE_0 struct {base.NoOperandsInstruction}
func (self *LSTORE_0) Execute(frame *runtime.Frame) {
	_lstore(frame,0)
}
type LSTORE_1 struct {base.NoOperandsInstruction}
func (self *LSTORE_1) Execute(frame *runtime.Frame) {
	_lstore(frame,1)
}
type LSTORE_2 struct {base.NoOperandsInstruction}
func (self *LSTORE_2) Execute(frame *runtime.Frame) {
	_lstore(frame,2)
}
type LSTORE_3 struct {base.NoOperandsInstruction}
func (self *LSTORE_3) Execute(frame *runtime.Frame) {
	_lstore(frame,3)
}
