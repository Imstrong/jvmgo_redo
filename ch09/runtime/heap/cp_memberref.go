package heap

import "jvmgo_redo/ch09/classfile"
//类成员标识符，包括成员变量和成员函数
type MemberRef struct {
	SymRef
	name string
	descriptor string
}
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className=refInfo.ClassName()
	self.name,self.descriptor=refInfo.NameAndDescriptor()
}
