package heap

import (
	"jvmgo/ch08/classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}
func newInterfaceMethodRef(cp *ConstantPool,refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref:=&InterfaceMethodRef{}
	ref.cp=cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method==nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}
func (self *InterfaceMethodRef) resolveInterfaceMethodRef()  {
	d:=self.cp.class//取得class信息
	c:=self.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method:=lookupInterfaceMethod(c,self.name,self.descriptor)
	if method==nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method=method
}
func lookupInterfaceMethod(c *Class,methodName,descriptor string) *Method{
	for _,method:=range c.methods {
		if methodName==method.Name()&&descriptor==method.Descriptor() {
			return method
		}
	}
	method:= lookupMethodInInterfaces(c.interfaces,methodName,descriptor)
	if method!=nil {
		return method
	}
	return nil
}