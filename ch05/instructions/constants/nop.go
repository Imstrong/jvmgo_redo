package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
