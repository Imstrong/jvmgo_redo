package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/runtime"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }
// hack!
func (self *INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopRef()
}
