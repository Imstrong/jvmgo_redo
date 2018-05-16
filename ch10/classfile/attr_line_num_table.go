package classfile
//指示java源码行号与jvm字节码的对应关系
type LineNumTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}
func (self *LineNumTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength:=reader.readUint16()
	self.lineNumberTable=make([]*LineNumberTableEntry,lineNumberTableLength)
	for i:=range self.lineNumberTable {
		self.lineNumberTable[i]=&LineNumberTableEntry{
			startPc:reader.readUint16(),
			lineNumber:reader.readUint16(),
		}
	}
}
func (self *LineNumTableAttribute) GetLineNumber(pc int) int {
	for i:=len(self.lineNumberTable)-1;i>=0;i-- {
		entry:=self.lineNumberTable[i]
		if pc>=int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}