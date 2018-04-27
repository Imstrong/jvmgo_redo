package heap

import "jvmgo_redo/ch05/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
}
func newMethods(class *Class,methodMembers []*classfile.AttrMethodInfo) []*Method{
	methods:=make([]*Method,len(methodMembers))
	for i,method:=range methodMembers {
		methods[i]=&Method{}
		methods[i].class=class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
	}
	return methods
}
func (self *Method) copyAttributes(methodMember *classfile.AttrMethodInfo) {
	if codeAttr:=methodMember.CodeAttribute();codeAttr!=nil {
		self.maxStack=codeAttr.MaxStack()
		self.maxLocals=codeAttr.MaxLocals()
		self.code=codeAttr.Code()
	}
}
