package conversions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/runtime"
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
type I2B struct {base.NoOperandsInstruction}
func (self *I2B) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	b:=int32(uint16(i))
	stack.PushInt(b)
}
type I2C struct {base.NoOperandsInstruction}
func (self *I2C) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	c:=int32(uint16(i))
	stack.PushInt(c)
}
type I2S struct {base.NoOperandsInstruction}
func (self *I2S) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	i:=stack.PopInt()
	s:=int32(uint16(i))
	stack.PushInt(s)
}