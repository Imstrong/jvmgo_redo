package heap

import (
	"jvmgo/ch09/classfile"
)

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, methodMembers []*classfile.AttrMethodInfo) []*Method {
	methods := make([]*Method, len(methodMembers))
	for i, method := range methodMembers {
		methods[i] = newMethod(class,method)
		//methods[i].class = class
		//methods[i].copyMemberInfo(method)
		//methods[i].copyAttributes(method)
		//methods[i].countArgSlotCount()
	}
	return methods
}
func (self *Method) copyAttributes(methodMember *classfile.AttrMethodInfo) {
	if codeAttr := methodMember.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
func (self *Method) Class() *Class {
	return self.class
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) IsStatic() bool {
	if self.accessFlags&ACC_STATIC != 0 {
		return true
	}
	return false
}
func (self *Method) Name() string {
	return self.name
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
func (self *Method) countArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
		if !self.IsStatic() {
			self.argSlotCount++
		}
	}
}
func (self *Method) IsAbstract() bool {
	if self.accessFlags&ACC_ABSTRACT != 0 {
		return true
	}
	return false
}
func (self *Method) IsNative() bool {
	if self.accessFlags&ACC_NATIVE != 0 {
		return true
	}
	return false
}
func newMethod(class *Class,methodAttr *classfile.AttrMethodInfo) *Method {
	method:=&Method{}
	method.class=class
	method.copyMemberInfo(methodAttr)
	method.copyAttributes(methodAttr)
	md:=parseMethodDescriptor(method.descriptor)
	method.countArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}
func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack=4
	self.maxLocals=self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code=[]byte{0xfe,0xb1}
	case 'D':
		self.code=[]byte{0xfe,0xaf}
	case 'F':
		self.code=[]byte{0xfe,0xae}
	case 'J':
		self.code=[]byte{0xfe,0xad}
	case 'L','[':
		self.code=[]byte{0xfe,0xb0}
	default:
		self.code=[]byte{0xfe,0xac}
	}
}