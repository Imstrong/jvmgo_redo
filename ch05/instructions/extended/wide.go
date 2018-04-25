package extended

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch05/instructions/loads"
	"jvmgo_redo/ch05/instructions/math"
	"jvmgo_redo/ch05/runtime"
	"jvmgo_redo/ch05/instructions/stores"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		//iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst // iload
	case 0x16: 
		// lload
		inst:=&loads.LLOAD{}
		inst.Index=uint(reader.ReadUint16())
		self.modifiedInstruction=inst
	case 0x17:
		// fload
		inst:=&loads.FLOAD{}
		inst.Index=uint(reader.ReadUint16())
		self.modifiedInstruction=inst
	case 0x18:
		// dload
		inst:=&loads.DLOAD{}
		inst.Index=uint(reader.ReadUint16())
		self.modifiedInstruction=inst
	case 0x19:
		// aload
		inst:=&loads.ALOAD{}
		inst.Index=uint(reader.ReadUint16())
		self.modifiedInstruction=inst
	case 0x36:
		// istore
		inst:=&stores.ISTORE{}
		i:=uint(reader.ReadUint8())
		inst.Index8Instruction=base.Index8Instruction{Index:i}
		self.modifiedInstruction=inst
	case 0x37:
		// lstore
		inst:=&stores.LSTORE{}
		i:=uint(reader.ReadUint8())
		inst.Index8Instruction=base.Index8Instruction{Index:i}
		self.modifiedInstruction=inst
	case 0x38:
		// fstore
		inst:=&stores.FSTORE{}
		i:=uint(reader.ReadUint8())
		inst.Index8Instruction=base.Index8Instruction{Index:i}
		self.modifiedInstruction=inst
	case 0x39:
		// dstore
		inst:=&stores.DSTORE{}
		i:=uint(reader.ReadUint8())
		inst.Index8Instruction=base.Index8Instruction{Index:i}
		self.modifiedInstruction=inst

	case 0x3a:
		// astore
		inst:=&stores.ASTORE{}
		i:=uint(reader.ReadUint8())
		inst.Index8Instruction=base.Index8Instruction{Index:i}
		self.modifiedInstruction=inst
	case 0x84:
		//IINC指令给局部变量表的索引增加常量值，索引和常量值都由操作数提供
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