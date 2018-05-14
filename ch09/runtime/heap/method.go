package heap

import "jvmgo_redo/ch08/classfile"

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
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
		methods[i].countArgSlotCount()
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
func (self *Method) countArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
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
