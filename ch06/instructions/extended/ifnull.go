package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/runtime"
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
