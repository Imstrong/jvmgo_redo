package classfile
//被final定义的常量值
type ConstantValueAttr struct {
	constantValueIndex uint16
}
func (self *ConstantValueAttr) readInfo(reader *ClassReader) {
	self.constantValueIndex=reader.readUint16()
}
func (self *ConstantValueAttr) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
