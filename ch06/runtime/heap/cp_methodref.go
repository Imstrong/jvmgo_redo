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

