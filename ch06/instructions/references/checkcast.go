package references

import (
	"jvmgo_redo/ch06/instructions/base"
	"jvmgo_redo/ch06/runtime"
)

type CHECK_CAST struct {base.Index16Instruction}
func (self *CHECK_CAST) Execute(frame *runtime.Frame) {

}
