package conversions

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

//convert float32 to x
type F2I struct {base.NoOperandsInstruction}
func (self *F2I) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	f:=stack.PopFloat()
	i:=int32(f)
	stack.PushInt(i)
}
type F2D struct {base.NoOperandsInstruction}
func(self *F2D) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	f:=stack.PopFloat()
	d:=float64(f)
	stack.PushDouble(d)
}
type F2L struct {base.NoOperandsInstruction}
func (self *F2L) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	f:=stack.PopFloat()
	l:=int64(f)
	stack.PushLong(l)
}
