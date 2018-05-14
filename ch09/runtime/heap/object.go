package heap
//java对象
type Object struct{
	//类信息
	class *Class
	//实例变量
	//fields Slots
	//用interface类型，可以表示数组。对于普通对象，存放的仍是Slots，对于数组，可以存放各种类型
	data interface{}
	//用于指向类对象（Class类的对象）的指针
	extra interface{}
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
func (object *Object) Data() interface{} {
	return object.data
}
func (object *Object) IsInstanceof(class *Class) bool {
	return class.isAssignableFrom(object.class)
}
func (object *Object) Class() *Class {
	return object.class
}
func (object *Object) Extra() interface{} {
	return object.extra
}