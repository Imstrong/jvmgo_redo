package heap

import (
	"jvmgo/ch08/classpath"
	"fmt"
	"jvmgo/ch08/classfile"
)

type ClassLoader struct {
	classpath *classpath.Classpath
	classMap map[string]*Class
	verboseFlag bool
}
func NewClassLoader(cp *classpath.Classpath,verboseInstFlag bool) *ClassLoader {
	return &ClassLoader{
		classpath:cp,
		verboseFlag:verboseInstFlag,
		classMap:make(map[string]*Class),
	}
}
func (self *ClassLoader) LoadClass(name string) *Class {
	if class,ok:=self.classMap[name];ok{
		return class
	}
	if name[0]=='[' {
		return self.LoadArrayClass(name)
	}
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	//类加载3步骤，一：读取
	data,entry:=self.readClass(className)
	//二：解析
	class:=self.defineClass(data)
	//三：连接
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", className, entry)
	}
	return class
}
func (self *ClassLoader) LoadArrayClass(name string) *Class {

}
func (self *ClassLoader) readClass(className string) ([]byte,classpath.Entry ){
	data,entry,err:=self.classpath.ReadClass(className)
	if err!=nil {
		panic("java.lang.ClassNotFoundException:"+className)
	}
	return data,entry
}
func (self *ClassLoader) defineClass(data []byte) *Class {
	//从字节数据读取并封装成结构体
	class:=parseClass(data)
	class.loader=self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name]=class
	return class
}
func parseClass(data []byte) *Class {
	cf,err:=classfile.Parse(data)
	if err!=nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
func resolveSuperClass(class *Class) {
	if class.name!="java/lang/Object" {
		//除了java.lang.Object所有的类都有且仅有一个超类
		class.superClass=class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount:=len(class.interfaces)
	if interfaceCount>0 {
		class.interfaces=make([]*Class,interfaceCount)
		for i,interfaceName:=range class.interfaceNames {
			class.interfaces[i]=class.loader.LoadClass(interfaceName)
		}
	}
}
func link(class *Class) {
	verify(class)
	prepare(class)
}
func verify(class *Class) {
	//empty
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
func calcInstanceFieldSlotIds(class *Class) {
	//计算实例字段
	slotId:=uint(0)
	//从父类开始算
	if class.superClass!=nil {
		slotId=class.superClass.instanceSlotCount
	}
	for _,field:=range class.fields {
		if !field.IsStatic() {
			field.slotId=slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount=slotId
}
func calcStaticFieldSlotIds(class *Class) {
	slotId:=uint(0)
	for _,field:=range class.fields {
		if field.IsStatic() {
			field.slotId=slotId
			slotId++
			if field.isLongOrDouble() {
				slotId--
			}
		}
	}
	class.staticSlotCount=slotId
}
func allocAndInitStaticVars(class *Class) {
	class.staticVars=newSlots(class.staticSlotCount)
	for _,field:=range class.fields {
		if field.IsStatic()&&field.IsFinal()  {
			initStaticFinalVar(class,field)
		}
	}
}
func initStaticFinalVar(class *Class,field *Field) {
	vars:=class.staticVars
	cp:=class.constantPool
	cpIndex:=field.ConstantValueIndex()
	slotId:=field.slotId
	if cpIndex>0 {
		switch field.Descriptor() {
		case "Z","B","C","S","I":
			val:=cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId,val)
		case "J":
			val:=cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId,val)
		case "F":
			val:=cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId,val)
		case "D":
			val:=cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId,val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}