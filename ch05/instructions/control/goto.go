package control

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch04/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
