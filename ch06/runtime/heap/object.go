package heap
type Object struct{
	//类信息
	class *Class
	//实例变量
	fields Slots
}
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}