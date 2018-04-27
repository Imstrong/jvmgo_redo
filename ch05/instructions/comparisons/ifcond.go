package comparisons

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch05/runtime"
)
//和0做比较
//if equals to zero
type IFEQ struct{ base.BranchInstruction }
func (self *IFEQ) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
//if not equals to zero
type IFNE struct{ base.BranchInstruction }
func (self *IFNE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}
//if lower than zero
type IFLT struct{ base.BranchInstruction }
func (self *IFLT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}
//if lower than or equals to zero
type IFLE struct{ base.BranchInstruction }
func (self *IFLE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}
//if greater than zero
type IFGT struct{ base.BranchInstruction }
func (self *IFGT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}
//if greater than or equals to zero
type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
