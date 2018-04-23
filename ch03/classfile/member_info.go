package classfile
type AttrMethodInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}
func readMembers(reader *ClassReader,cp ConstantPool) []*AttrMethodInfo {
	memberCount:=reader.readUint16()
	members:=make([]*AttrMethodInfo,memberCount)
	for i:=range members {
		members[i]=readMember(reader,cp)
	}
	return members
}
func readMember(reader *ClassReader,cp ConstantPool) *AttrMethodInfo {
	accessFlags:=reader.readUint16()
	nameIndex:=reader.readUint16()
	descriptorIndex:=reader.readUint16()
	attributes:=readAttributes(reader,cp)
	return &AttrMethodInfo{
		cp:cp,
		accessFlags:accessFlags,
		nameIndex:nameIndex,
		descriptorIndex:descriptorIndex,
		attributes:attributes,
	}
}
func (self *AttrMethodInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *AttrMethodInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

