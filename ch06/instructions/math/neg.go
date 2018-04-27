package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/runtime"
)

type INEG struct {
	base.NoOperandsInstruction
}
type LNEG struct {
	base.NoOperandsInstruction
}
type FNEG struct {
	base.NoOperandsInstruction
}
type DNEG struct {
	base.NoOperandsInstruction
}
func (self *INEG) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val:=-stack.PopInt()
	stack.PushInt(val)
}
func (self *LNEG) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val:=-stack.PopLong()
	stack.PushLong(val)
}
func (self *FNEG) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val:=-stack.PopFloat()
	stack.PushFloat(val)
}
func (self *DNEG) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val:=-stack.PopDouble()
	stack.PushDouble(val)
}
