package base

import "jvmgo_redo/ch06/runtime"

//指令接口，规定字节码指令的基本行为
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtime.Frame)
}
type NoOperandsInstruction struct {}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

//存放跳转指令，offset为偏移量
type BranchInstruction struct {
	Offset int
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset=int(reader.ReadInt16())
}

//根据索引存取局部变量表，索引为单字节操作数，index表示局部变量表索引
type Index8Instruction struct {
	Index uint
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index=uint(reader.ReadUint8())
}

//需要访问常量池的指令，常量池索引为2字节操作数
type Index16Instruction struct {
	Index uint
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index=uint(reader.ReadUint16())
}
