package math

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

type IDIV struct {
	base.NoOperandsInstruction
}
type LDIV struct {
	base.NoOperandsInstruction
}
type FDIV struct {
	base.NoOperandsInstruction
}
type DDIV struct {
	base.NoOperandsInstruction
}
func (self *IDIV) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	val:=v1/v2
	if v2==0 {
		panic("ArithmeticException: / by zero")
	}
	stack.PushInt(val)
}
func (self *LDIV) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	val:=v1/v2
	if v2==0 {
		panic("ArithmeticException: / by zero")
	}
	stack.PushLong(val)
}
func (self *FDIV) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	val:=v1/v2
	stack.PushFloat(val)
}
func (self *DDIV) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
	val:=v1/v2
	stack.PushDouble(val)
}
