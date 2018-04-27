package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}
func (self *ClassMember) copyMemberInfo (memberInfo *classfile.AttrMethodInfo) {
	self.accessFlags=memberInfo.AccessFlags()
	self.name=memberInfo.Name()
	self.descriptor=memberInfo.Descriptor()
}
