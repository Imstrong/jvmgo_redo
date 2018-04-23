package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

type IF_ACMPEQ struct {base.BranchInstruction}
func (self *IF_ACMPEQ) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	ref2:=stack.PopRef()
	ref1:=stack.PopRef()
	if ref1==ref2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ACMPNE struct {base.BranchInstruction}
