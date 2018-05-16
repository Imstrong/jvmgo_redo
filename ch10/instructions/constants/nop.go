package constants

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
