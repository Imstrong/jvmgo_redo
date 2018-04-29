package heap

import "math"

//Slots表示局部变量表，局部变量可以为数字类型，因此会有一个属性是空的（0）
//go不允许包互相依赖，即a依赖b的同时b依赖a
type Slot struct {
	num int32
	ref *Object
}
type Slots []Slot
func newSlots (maxLocals uint) Slots {
	if maxLocals>0 {
		return make([]Slot,maxLocals)
	}
	return nil
}
//int
func (self Slots) SetInt(index uint,val int32) {
	self[index].num=val
}
func (self Slots) GetInt(index uint) int32{
	return self[index].num
}
//float
func (self Slots) SetFloat(index uint,val float32) {
	bits:=math.Float32bits(val)
	self[index].num=int32(bits)
}
func (self Slots) GetFloat (index uint) float32{
	bits:=uint32(self[index].num)
	return math.Float32frombits(bits)
}
//long占两个位置
func (self Slots) SetLong(index uint,val int64) {
	self[index].num=int32(val)
	self[index+1].num=int32(val>>32)
}
func (self Slots) GetLong(index uint) int64{
	low:=self[index].num
	hign:=self[index+1].num
	return int64(hign)<<32|int64(low)
}
//double
func(self Slots) SetDouble(index uint,val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index,int64(bits))
}
func(self Slots) GetDouble(index uint) float64 {
	val := uint64(self.GetLong(index))
	return math.Float64frombits(val)
}
//ref
func (self Slots) SetRef(index uint,ref *Object) {
	self[index].ref=ref
}
func(self Slots) GetRef(index uint) *Object {
	return self[index].ref
}

