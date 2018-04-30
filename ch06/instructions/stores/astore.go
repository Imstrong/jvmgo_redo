package stores

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)

//存储指令从操作数栈弹出变量，并存入局部变量表
func _astore(frame *runtime.Frame,index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index,val)
}
type ASTORE struct {base.Index8Instruction}
func (self *ASTORE) Execute(frame *runtime.Frame) {
	_astore(frame,self.Index)
}
type ASTORE_0 struct {base.NoOperandsInstruction}
func (self *ASTORE_0) Execute(frame *runtime.Frame) {
	_astore(frame,0)
}
type ASTORE_1 struct {base.NoOperandsInstruction}
func (self *ASTORE_1) Execute(frame *runtime.Frame) {
	_astore(frame,1)
}
type ASTORE_2 struct {base.NoOperandsInstruction}
func (self *ASTORE_2) Execute(frame *runtime.Frame) {
	_astore(frame,2)
}
type ASTORE_3 struct {base.NoOperandsInstruction}
func (self *ASTORE_3) Execute(frame *runtime.Frame) {
	_astore(frame,3)
}
