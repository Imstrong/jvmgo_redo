package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

type DUP struct {base.NoOperandsInstruction}
func (self *DUP) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP_X1 struct {base.NoOperandsInstruction}
func (self *DUP_X1) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP_X2 struct {base.NoOperandsInstruction}
func (self *DUP_X2) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP2 struct {base.NoOperandsInstruction}
func (self *DUP2) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP2_X1 struct {base.NoOperandsInstruction}
func (self *DUP2_X1) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
type DUP2_X2 struct {base.NoOperandsInstruction}
func (self *DUP2_X2) Execute (frame *runtime.Frame) {
	stack:=frame.OperandStack()
	slot:=stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
