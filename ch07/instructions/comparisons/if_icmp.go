package comparisons

import (
	"jvmgo_redo/ch07/instructions/base"
	"jvmgo_redo/ch07/runtime"
)
//判断两个int的大小
type IF_ICMPEQ struct {base.BranchInstruction}
func (self *IF_ICMPEQ) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1==val2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ICMPNE struct {base.BranchInstruction}
func (self *IF_ICMPNE) Execute(frame *runtime.Frame ) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1!=val2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ICMPGT struct {base.BranchInstruction}
func (self *IF_ICMPGT) Execute(frame *runtime.Frame ) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1>val2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ICMPGE struct {base.BranchInstruction}
func (self *IF_ICMPGE) Execute(frame *runtime.Frame ) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1>=val2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ICMPLT struct {base.BranchInstruction}
func (self *IF_ICMPLT) Execute(frame *runtime.Frame ) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1<val2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ICMPLE struct {base.BranchInstruction}
func (self *IF_ICMPLE) Execute(frame *runtime.Frame ) {
	stack:=frame.OperandStack()
	val2:=stack.PopInt()
	val1:=stack.PopInt()
	if val1<=val2 {
		base.Branch(frame,self.Offset)
	}
}
