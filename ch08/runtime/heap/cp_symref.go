package heap
//symbol reference 符号引用
type SymRef struct {
	//常量池指针，通过常量池又可以访问类数据
	cp *ConstantPool
	//类名
	className string
	//类结构体指针，通过常量池获取类结构体地址后缓存在这个字段
	class *Class
}
func (self *SymRef) ResolvedClass() *Class {
	if self.class==nil {
		self.resolveClassRef()
	}
	return self.class
}
func (self *SymRef) resolveClassRef() {
	//通过常量池指针获得类结构体指针
	d:=self.cp.class
	c:=d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.class=c
}
