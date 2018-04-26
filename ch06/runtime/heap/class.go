package heap

import (
	"jvmgo/ch05/classfile"
)

type Class struct {
	//类访问标识
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantPool *classfile.ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	staticVars *Slots
}
