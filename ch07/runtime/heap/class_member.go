package heap

import "jvmgo_redo/ch07/classfile"

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
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c:=self.Class()
	if self.IsProtected() {
		return d==c||d.isSubClassOf(c)||
			c.getPackageName()==d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName()==d.getPackageName()
	}
	b:= d==c
	return b
}
func (self *ClassMember) IsPublic() bool {
	if self.accessFlags&ACC_PUBLIC!=0 {
		return true
	}
	return false
}
func (self *ClassMember) IsProtected() bool {
	if self.accessFlags&ACC_PROTECTED!=0 {
		return true
	}
	return false
}
func (self *ClassMember) IsPrivate() bool {
	if self.accessFlags&ACC_PRIVATE!=0 {
		return true
	}
	return false
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}