package heap
type Object struct{
	//类信息
	class *Class
	//实例变量
	fields Slots
}
