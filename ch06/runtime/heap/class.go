package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	//类访问标识
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	//类加载器指针
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
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
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic()||self.getPackageName()==other.getPackageName()
}
func (self *Class) getPackageName() string {
	if i:=strings.LastIndex(self.name,"/");i>=0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) isSubClassOf(d *Class) bool {

}
func (self *Class) NewObject() *Object {
	return newObject(self)
}
func newObject(class *Class) *Object {
	return &Object{
		class:class,
		fields:newSlots(class.instanceSlotCount),
	}
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}