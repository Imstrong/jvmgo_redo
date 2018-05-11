package runtime

import (
	"math"
	"jvmgo/ch08/runtime/heap"
)

type LocalVars []Slot
func newLocalVars (maxLocals uint) LocalVars {
	if maxLocals>0 {
		return make([]Slot,maxLocals)
	}
	return nil
}
//int
func (self LocalVars) SetInt(index uint,val int32) {
	self[index].num=val
}
func (self LocalVars) GetInt(index uint) int32{
	return self[index].num
}
//float
func (self LocalVars) SetFloat(index uint,val float32) {
	bits:=math.Float32bits(val)
	self[index].num=int32(bits)
}
func (self LocalVars) GetFloat (index uint) float32{
	bits:=uint32(self[index].num)
	return math.Float32frombits(bits)
}
//long占两个位置
func (self LocalVars) SetLong(index uint,val int64) {
	self[index].num=int32(val)
	self[index+1].num=int32(val>>32)
}
func (self LocalVars) GetLong(index uint) int64{
	low:=self[index].num
	hign:=self[index+1].num
	return int64(hign)<<32|int64(low)
}
//double
func(self LocalVars) SetDouble(index uint,val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index,int64(bits))
}
func(self LocalVars) GetDouble(index uint) float64 {
	val := uint64(self.GetLong(index))
	return math.Float64frombits(val)
}
//ref
func (self LocalVars) SetRef(index uint,ref *heap.Object) {
	self[index].ref=ref
}
func(self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
func (self LocalVars) SetSlot(index uint,slot Slot) {
	self[index]=slot
}