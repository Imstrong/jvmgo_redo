package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
	class *Class
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
func (self *Method) Class() *Class {
	return self.class
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) MaxStack() uint{
	return self.maxStack
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) IsStatic() bool {
	if self.accessFlags&ACC_STATIC!=0 {
		return true
	}
	return false
}
func (self *Method) Name() string {
	return self.name
}