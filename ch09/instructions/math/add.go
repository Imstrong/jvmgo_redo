package math

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
)

type IADD struct {
	base.NoOperandsInstruction
}
type LADD struct {
	base.NoOperandsInstruction
}
type FADD struct {
	base.NoOperandsInstruction
}
type DADD struct {
	base.NoOperandsInstruction
}
func (self *IADD) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	val:=v1+v2
	stack.PushInt(val)
}
func (self *LADD) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	val:=v1+v2
	stack.PushLong(val)
}
func (self *FADD) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	val:=v1+v2
	stack.PushFloat(val)
}
func (self *DADD) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	val:=v1+v2
	stack.PushDouble(val)
}
