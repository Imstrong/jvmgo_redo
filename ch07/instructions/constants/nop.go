package constants

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
