package heap

import "jvmgo_redo/ch06/classfile"

type MemberRef struct {
	SymRef
	name string
	descriptor string
}
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className=refInfo.ClassName()
	self.name,self.descriptor=refInfo.NameAndDescriptor()
}
