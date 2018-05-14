package heap

import "jvmgo/ch09/classfile"
//成员函数表示符
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
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method==nil {
		self.resolveMethodRef()
	}
	return self.method
}
func (self *MethodRef) resolveMethodRef() {
	d:=self.cp.class
	c:=self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassException")
	}
	method:=lookupMethod(c,self.name,self.descriptor)
	if method==nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessException")
	}
	self.method=method
}
func lookupMethod(c *Class,name,descriptor string) *Method{
	method:=LookupMethodInClass(c,name,descriptor)
	if method==nil {
		method=lookupMethodInInterfaces(c.interfaces,name,descriptor)
	}
	return method
}
func LookupMethodInClass(class *Class,name,descriptor string) *Method {
	//c=class执行一次，如果在本类中找到匹配的方法，返回；否则在父类中查找
	for c:=class;c!=nil;c=c.superClass {
		for _,method:=range c.methods {
			if method.name==name&&method.descriptor==descriptor {
				return method
			}
		}
	}
	return nil
}
func lookupMethodInInterfaces(interfaces []*Class,name,descriptor string) *Method {
	for _,iface:=range interfaces {
		for _,method:=range iface.methods {
			if method.name==name && method.descriptor==descriptor {
				return method
			}
			method:=lookupMethodInInterfaces(iface.interfaces,name,descriptor)
			if method!=nil {
				return method
			}
		}
	}
	return nil
}
