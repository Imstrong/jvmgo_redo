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
func (object *Object) SetRefVar(refName,descriptor string,refVal *Object){
	field:=object.class.getField(refName,descriptor,false)
	slots:=object.data.(Slots)
	slots.SetRef(field.slotId,refVal)
}
func (object *Object) GetRefVar(name,descriptor string) *Object {
	field:=object.Class().getField(name,descriptor,false)
	slots:=object.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (object *Object) Clone() *Object {
	return &Object{
		class:object.class,
		data:object.cloneData(),
	}
}
func (object *Object) cloneData() interface{} {
	switch object.data.(type) {
	case []int8:
		elements:=object.data.([]int8)
		elements2:=make([]int8,len(elements))
		copy(elements2,elements)
		return elements2
	case []int16:
		elements:=object.data.([]int16)
		elements2:=make([]int16,len(elements))
		copy(elements2,elements)
		return elements2
	case []int32:
		elements:=object.data.([]int32)
		elements2:=make([]int32,len(elements))
		copy(elements2,elements)
		return elements2
	case []int64:
		elements:=object.data.([]int64)
		elements2:=make([]int64,len(elements))
		copy(elements2,elements)
		return elements2
	case []uint16:
		elements:=object.data.([]uint16)
		elements2:=make([]uint16,len(elements))
		copy(elements2,elements)
		return elements2
	case []float32:
		elements:=object.data.([]float32)
		elements2:=make([]float32,len(elements))
		copy(elements2,elements)
		return elements2
	case []float64:
		elements:=object.data.([]float64)
		elements2:=make([]float64,len(elements))
		copy(elements2,elements)
		return elements2
	case []*Object:
		elements:=object.data.([]*Object)
		elements2:=make([]*Object,len(elements))
		copy(elements2,elements)
		return elements2
	default:
		slots:=object.data.(Slots)
		slots2:=newSlots(uint(len(slots)))
		copy(slots2,slots)
		return slots2
	}
}