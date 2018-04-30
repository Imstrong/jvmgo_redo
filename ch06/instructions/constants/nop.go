package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
