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
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c:=self.class
	if self.IsProtected() {
		return d==c||d.isSubClassOf(c)||
			c.getPackageName()==d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName()==d.getPackageName()
	}
	return d==c
}
func (self *ClassMember) IsPublic() bool {

}
func (self *ClassMember) IsProtected() bool {

}
func (self *ClassMember) IsPrivate() bool {

}