package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
)

//no operation
type NOP struct {
	base.NoOperandsInstruction
}
func (self *NOP) Execute(frame *runtime.Frame) {
	//do nothing
}
