package extended

import (
	"jvmgo_redo/ch08/instructions/base"
	"jvmgo_redo/ch08/runtime"
)

//根据引用该是否null进行跳转
type IFNULL struct {base.BranchInstruction}
type IFNONNULL struct {base.BranchInstruction}
func (self *IFNULL) Execute(frame *runtime.Frame) {
	ref:=frame.OperandStack().PopRef()
	if ref==nil {
		base.Branch(frame,self.Offset)
	}
}
func (self *IFNONNULL) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref!=nil {
		base.Branch(frame,self.Offset)
	}
}
