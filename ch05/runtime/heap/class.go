package heap

import (
	"jvmgo_redo/ch05/classfile"
)

type Class struct {
	//类访问标识
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *classfile.ConstantPool
	fields            []*Field
	methods           []*Method
	//类加载器指针
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}

func newClass(classfile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classfile.AccessFlags()
	class.name = classfile.ClassName()
	class.superClassName = classfile.SuperClassName()
	class.interfaceNames = classfile.InterfaceNames()
	class.constantPool = newConstantPool(class, classfile.ConstantPool())
	class.fields = newFields(class, classfile.Fields())
	class.methods = newMethods(class, classfile.Methods())
	return class
}

//判断访问权限标识
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
