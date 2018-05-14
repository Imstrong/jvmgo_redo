package heap

import "jvmgo_redo/ch08/classfile"

//类符号引用，继承符号引用结构
type ClassRef struct {
	SymRef
}
func newClassRef(cp *ConstantPool,classInfo *classfile.ConstantClassInfo) *ClassRef{
	ref:=&ClassRef{}
	ref.cp=cp
	ref.className=classInfo.Name()
	return ref
}

