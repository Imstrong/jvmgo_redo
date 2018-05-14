package control

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
