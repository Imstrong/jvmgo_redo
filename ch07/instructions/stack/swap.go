package stack

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)
//将栈顶两个变量交换位置
type SWAP struct {
	base.NoOperandsInstruction
}
func (self *SWAP) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot1:=stack.PopSlot()
	slot2:=stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
