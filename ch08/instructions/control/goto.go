package control

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

type GOTO struct {
	base.BranchInstruction
}
func (self *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.Offset)
}
