package stores

import (
	"jvmgo_redo/ch08/instructions/base"
	"jvmgo_redo/ch08/runtime"
)

//存储指令从操作数栈弹出变量，并存入局部变量表
func _dstore(frame *runtime.Frame,index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index,val)
}
type DSTORE struct {base.Index8Instruction}
func (self *DSTORE) Execute(frame *runtime.Frame) {
	_dstore(frame,self.Index)
}
type DSTORE_0 struct {base.NoOperandsInstruction}
func (self *DSTORE_0) Execute(frame *runtime.Frame) {
	_dstore(frame,0)
}
type DSTORE_1 struct {base.NoOperandsInstruction}
func (self *DSTORE_1) Execute(frame *runtime.Frame) {
	_dstore(frame,1)
}
type DSTORE_2 struct {base.NoOperandsInstruction}
func (self *DSTORE_2) Execute(frame *runtime.Frame) {
	_dstore(frame,2)
}
type DSTORE_3 struct {base.NoOperandsInstruction}
func (self *DSTORE_3) Execute(frame *runtime.Frame) {
	_dstore(frame,3)
}
