package math

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch04/runtime"
)

//给局部变量表中的int变量增加常量值，局部变量表的索引和常量值都由操作数提供
type IINC struct {
	Index uint
	Const int32
}
//从字节码中读取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index=uint(reader.ReadUint8())
	self.Const=int32(reader.ReadInt8())
}
//从局部变量表读取变量，加上常量值，再把结果写回局部变量表
func (self *IINC) Execute(frame *runtime.Frame) {
	localVars:=frame.LocalVars()
	val:=localVars.GetInt(self.Index)
	localVars.SetInt(self.Index,val)
}



