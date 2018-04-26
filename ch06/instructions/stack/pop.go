package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

type POP struct {base.NoOperandsInstruction}
func (self *POP) Execute(frame *runtime.Frame ){
	stack:=frame.OperandStack()
	stack.PopSlot()
}
type POP2 struct {base.NoOperandsInstruction}
func (self *POP2) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
