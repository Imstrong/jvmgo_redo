package extended

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

type GOTO_W struct {
	offset int
}
func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset=int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *runtime.Frame) {
	base.Branch(frame,self.offset)
}
