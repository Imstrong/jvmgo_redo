package classfile
//方法表：java代码编译成的字节码指令
type CodeAttribute struct  {
	cp ConstantPool
	maxStack uint16
	maxLocals uint16
	code []byte
	exceptionTable []*ExceptionTableEntry
	attributes []AttributeInfo
}
type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack=reader.readUint16()
	self.maxLocals=reader.readUint16()
	codeLength:=reader.readUint32()
	self.code=reader.readBytes(codeLength)
	self.exceptionTable=readExceptionTable(reader)
	self.attributes=readAttributes(reader,self.cp)
}
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptiontableLength:=reader.readUint16()
	exceptionTable:=make([]*ExceptionTableEntry,exceptiontableLength)
	for i:= range exceptionTable {
		exceptionTable[i]=&ExceptionTableEntry{
			startPc:reader.readUint16(),
			endPc:reader.readUint16(),
			handlerPc:reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}