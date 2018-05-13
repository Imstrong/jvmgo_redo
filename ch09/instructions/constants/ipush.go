package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
)
//ipush指令将从操作数中获取一个byte、short型的整数，扩展成int型并推入栈顶
type BIPUSH struct {
	val int8
}
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val=reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *runtime.Frame) {
	i:=int32(self.val)
	frame.OperandStack().PushInt(i)
}
type SIPUSH struct {
	val int16
}
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val=reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *runtime.Frame){
	frame.OperandStack().PushInt(int32(self.val))
}

