package control

import (
	"jvmgo_redo/ch08/instructions/base"
	"jvmgo_redo/ch08/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
