package math

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
)

//布尔运算指令，有按位与、按位或、按位异或三种，且只能操作int、long两种类型，因此组合成六种
//按位与
type IAND struct {
	base.NoOperandsInstruction
}
func (self *IAND) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	result:=v1&v2
	stack.PushInt(result)
}
type LAND struct {
	base.NoOperandsInstruction
}
func (self *LAND) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	result:=v1&v2
	stack.PushLong(result)
}
type IOR struct {base.NoOperandsInstruction}
func (self *IOR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	result:=v1|v2
	stack.PushInt(result)
}
type LOR struct{base.NoOperandsInstruction}
func(self *LOR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	result:=v1|v2
	stack.PushLong(result)
}
type IXOR struct{base.NoOperandsInstruction}
func(self *IXOR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	result:=v1^v2
	stack.PushInt(result)
}
type LXOR struct{base.NoOperandsInstruction}
func(self *LXOR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	result:=v1^v2
	stack.PushLong(result)
}

