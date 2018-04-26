package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
