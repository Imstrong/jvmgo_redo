package constants

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
