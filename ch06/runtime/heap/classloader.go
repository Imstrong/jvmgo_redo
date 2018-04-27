package heap

import (
	"jvmgo/ch02/classpath"
	"fmt"
)

type ClassLoader struct {
	classpath *classpath.Classpath
	classMap map[string]*Class
}
func (self *ClassLoader) NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		classpath:cp,
		classMap:make(map[string]*Class),
	}
}
func (self *ClassLoader) loadClass(name string) *Class {
	if class,ok:=self.classMap[name];ok{
		return class
	}
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	data,entry:=self.readClass(className)
	class:=self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n",className,entry)
	return class
}
func (self *ClassLoader) readClass(className string) ([]byte,classpath.Entry ){
	data,entry,err:=self.classpath.ReadClass(className)
	if err!=nil {
		panic("java.lang.ClassNotFoundException:"+className)
	}
	return data,entry
}