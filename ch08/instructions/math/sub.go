package math

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

type ISUB struct {
	base.NoOperandsInstruction
}
type LSUB struct {
	base.NoOperandsInstruction
}
type FSUB struct {
	base.NoOperandsInstruction
}
type DSUB struct {
	base.NoOperandsInstruction
}
func (self *ISUB) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	val:=v1-v2
	stack.PushInt(val)
}
func (self *LSUB) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	val:=v1-v2
	stack.PushLong(val)
}
func (self *FSUB) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	val:=v1-v2
	stack.PushFloat(val)
}
func (self *DSUB) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	val:=v1-v2
	stack.PushDouble(val)
}
