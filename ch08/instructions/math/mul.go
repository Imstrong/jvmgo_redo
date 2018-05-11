package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
)

type IMUL struct {
	base.NoOperandsInstruction
}
type LMUL struct {
	base.NoOperandsInstruction
}
type FMUL struct {
	base.NoOperandsInstruction
}
type DMUL struct {
	base.NoOperandsInstruction
}
func (self *IMUL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	val:=v1*v2
	stack.PushInt(val)
}
func (self *LMUL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	val:=v1*v2
	stack.PushLong(val)
}
func (self *FMUL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	val:=v1*v2
	stack.PushFloat(val)
}
func (self *DMUL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	val:=v1*v2
	stack.PushDouble(val)
}
