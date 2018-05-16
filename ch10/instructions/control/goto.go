package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
