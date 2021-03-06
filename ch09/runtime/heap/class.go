package heap

import (
	"jvmgo_redo/ch09/classfile"
	"strings"
)

//加载后存放在方法区的类信息，以go数据结构的形式保存，不是java对象
type Class struct {
	//类访问标识
	accessFlags    uint16
	name           string
	superClassName string
	interfaceNames []string
	constantPool   *ConstantPool
	fields         []*Field
	methods        []*Method
	//类加载器指针
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	//Class实例
	jClass *Object
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
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) NewObject() *Object {
	return newObject(self)
}
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) IsSuperClassOf(class *Class) bool {
	if class.superClass == self {
		return true
	}
	return false
}
func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i > 0 {
		return self.name[:i]
	}
	return ""
}
func (self *Class) IsSubClassOf(class *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == class {
			return true
		}
	}
	return false
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) IsImplements(class *Class) bool {
	for i := self; i != nil; i = i.superClass {
		for _, iFace := range self.interfaces {
			if iFace == class || iFace.isSubInterfaceOf(class) {
				return true
			}
		}
	}
	return false
}
func (self *Class) Name() string {
	return self.name
}
func (class *Class) InitStarted() bool {
	return class.initStarted
}
func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}
func (self *Class) StartInit() {
	self.initStarted = true
}
func (self *Class) ClassLoader() *ClassLoader {
	return self.loader
}
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}
func (self *Class) JClass() *Object {
	return self.jClass
}
func (self *Class) IsPrimitive() bool {
	_,ok:=primitiveTypes[self.name]//如果 存在 这个基本类型
	return ok
}
func (self *Class) getField(name,descriptor string,isStatic bool) *Field {
	for c:=self;c!=nil;c=c.superClass {
		for _,field:=range c.fields {
			if field.IsStatic()==isStatic&&field.name==name&&field.descriptor==descriptor {
				return field
			}
		}
	}
	return nil
}
func (self *Class) GetRefVar(name,descriptor string ) *Object {
	field:=self.getField(name,descriptor,false)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) GetInstanceMethod(name,descriptor string) *Method {
	return self.GetMethod(name,descriptor,false)
}
func (self *Class) GetMethod(name,descriptor string,isStatic bool) *Method {
	for c:=self;c!=nil;c=c.superClass {
		for _,method:=range c.methods {
			if method.name==name&&method.descriptor==descriptor {
				return method
			}
		}
	}
	return nil
}