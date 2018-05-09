package comparisons

import (
	"jvmgo_redo/ch07/instructions/base"
	"jvmgo_redo/ch07/runtime"
)


//判断引用是否相等
type IF_ACMPEQ struct{base.BranchInstruction}
func (self *IF_ACMPEQ) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopRef()
	v1:=stack.PopRef()
	if v1==v2 {
		base.Branch(frame,self.Offset)
	}
}
type IF_ACMPNE struct {base.BranchInstruction}
func(self *IF_ACMPNE) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopRef()
	v1:=stack.PopRef()
	if v1!=v2 {
		base.Branch(frame,self.Offset)
	}
}
