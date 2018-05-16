package classfile
type ConstantStringInfo struct {
	constantPool ConstantPool
	stringIndex uint16
}
func (self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stringIndex=reader.readUint16()
}
func (self *ConstantStringInfo) String() string {
	return self.constantPool.getUtf8(self.stringIndex)
}

