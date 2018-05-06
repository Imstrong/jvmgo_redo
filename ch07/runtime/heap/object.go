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
func (object *Object) Fields() Slots {
	return object.fields
}
func (object *Object) IsInstanceof(class *Class) bool {
	return class.isAssignableFrom(object.class)
}
func (object *Object) Class() *Class {
	return object.class
}
