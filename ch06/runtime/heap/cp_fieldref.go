package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}
func newFiledRef(cp *ConstantPool,refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref:=&FieldRef{}
	ref.cp=cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

