package conversions

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

//convert long to x
type L2I struct{base.NoOperandsInstruction}
func (self *L2I) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	l:=stack.PopLong()
	i:=int32(l)
	stack.PushInt(i)
}
type L2F struct{base.NoOperandsInstruction}
func (self *L2F) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	l:=stack.PopLong()
	i:=float32(l)
	stack.PushFloat(i)
}
type L2D struct{base.NoOperandsInstruction}
func (self *L2D) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	l:=stack.PopLong()
	i:=float64(l)
	stack.PushDouble(i)
}
