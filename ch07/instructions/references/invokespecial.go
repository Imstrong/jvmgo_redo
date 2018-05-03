package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }
// hack!
func (self *INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopRef()
}
