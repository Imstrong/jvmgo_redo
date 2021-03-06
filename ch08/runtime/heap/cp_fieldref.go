package heap

import "jvmgo_redo/ch08/classfile"
//成员变量标识符
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
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}
func (self *FieldRef) resolveFieldRef() {
	d:=self.cp.class
	c:=self.ResolvedClass()
	field:=self.lookupField(c,self.name,self.descriptor)
	if field==nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field=field
}
func (self *FieldRef) lookupField(c *Class,name,descriptor string) *Field {
	for _,field:=range c.fields {
		if field.name==name&&field.descriptor==descriptor {
			return field
		}
	}
	for _,iface:=range c.interfaces {
		if field:=self.lookupField(iface,name,descriptor);field!=nil {
			return field
		}
	}
	if c.superClass!=nil {
		return self.lookupField(c.superClass,name,descriptor)
	}
	return nil
}
func (self *Field) IsFinal() bool {
	if self.accessFlags&ACC_FINAL!=0 {
		return true
	}
	return false
}