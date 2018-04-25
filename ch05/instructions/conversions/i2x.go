package conversions

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch04/runtime"
)

//convert int to others
type I2F struct {
	base.NoOperandsInstruction
}
func (self *I2F) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	f:=float32(i)
	stack.PushFloat(f)
}
type I2D struct {
	base.NoOperandsInstruction
}
func (self *I2D) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	d:=float64(i)
	stack.PushDouble(d)
}
type I2L struct {
	base.NoOperandsInstruction
}
func(self *I2L) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	l:=int64(i)
	stack.PushLong(l)
}