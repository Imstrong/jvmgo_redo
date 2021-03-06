package control

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
