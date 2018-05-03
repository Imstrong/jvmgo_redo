package heap

import "jvmgo_redo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}
func newMethodRef(cp *ConstantPool,refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref:=&MethodRef{}
	ref.cp=cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
func (self *MethodRef) Name() string {
	return self.name
}
func (self *MethodRef) Descriptor() string {
	return self.descriptor
}

