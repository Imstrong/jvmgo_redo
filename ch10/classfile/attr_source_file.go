package classfile
type SourcefileAttribute struct {
	cp ConstantPool
	sourcefileIndex uint16
}
func (self *SourcefileAttribute) readInfo(reader *ClassReader ) {
	self.sourcefileIndex=reader.readUint16()
}
func (self *SourcefileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourcefileIndex)
}
