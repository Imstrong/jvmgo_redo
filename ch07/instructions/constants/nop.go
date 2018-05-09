package constants

import (
	"jvmgo_redo/ch07/instructions/base"
	"jvmgo_redo/ch07/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
