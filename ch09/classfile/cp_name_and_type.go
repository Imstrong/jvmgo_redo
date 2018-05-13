package classfile
type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIncex uint16
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex=reader.readUint16()
	self.descriptorIncex=reader.readUint16()
}

