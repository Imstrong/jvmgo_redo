package extend

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch05/instructions/loads"
	"jvmgo_redo/ch05/instructions/math"
	"jvmgo_redo/ch05/runtime"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst // iload
	//case 0x16: ... // lload
	//case 0x17: ... // fload
	//case 0x18: ... // dload
	//case 0x19: ... // aload
	//case 0x36: ... // istore
	//case 0x37: ... // lstore
	//case 0x38: ... // fstore
	//case 0x39: ... // dstore
	//case 0x3a: ... // astore
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst // iinc
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}
func (self *WIDE) Execute(frame *runtime.Frame) {
	self.modifiedInstruction.Execute(frame)
}