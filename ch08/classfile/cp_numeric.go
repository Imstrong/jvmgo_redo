package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}
//ConstantIntegerInfo正好可以存储一个int类型变量，但实际上，boolean，byte，short，char类型的常量也放在ConstantIntegerInfo中
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes:=reader.readUint32()
	self.val=int32(bytes)
}
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}
type ConstantFloatInfo struct {
	val float32
}
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes:=reader.readUint32()
	self.val=math.Float32frombits(bytes)
}
type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes:=reader.readUint64()
	self.val=int64(bytes)
}
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}
type ConstantDoubleInfo struct {
	val float64
}
func(self ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes:=reader.readUint64()
	self.val=math.Float64frombits(bytes)
}
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}

