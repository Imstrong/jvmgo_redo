package conversions

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

//d2x:convert double to others,把dobule类型转换成其他类型
type D2F struct {base.NoOperandsInstruction}
func (self *D2F) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	i:=float32(d)
	stack.PushFloat(i)
}
type D2I struct {base.NoOperandsInstruction}
func (self *D2I) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	i:=int32(d)
	stack.PushInt(i)
}
type D2L struct {base.NoOperandsInstruction}
func (self *D2L) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	d:=stack.PopDouble()
	l:=int64(d)
	stack.PushLong(l)
}
