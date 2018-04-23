package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/runtime"
)

//位移指令，按int，long|left，right组合成六种情况
//逻辑右移就是无符号右移，java中对应运算符为>>>

//int左位移
type ISHL struct {
	base.NoOperandsInstruction
}
func (self *ISHL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopInt()
	//因为int只有32位，可以位移的位数用5位就可以表示（0-31），所以取低五位。因为操作符右侧必须为无符号整数，所以进行类型转换
	s:=uint32(v2)&0x1f
	//将操作数v1左移操作数v2位
	result:=v1<<s
	stack.PushInt(result)
}
//int算数右位移
type ISHR struct {
	base.NoOperandsInstruction
}
func (self ISHR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopInt()
	v1:=stack.PopLong()
	s:=uint32(v2)&0x3f
	result:=v1>>s
	stack.PushLong(result)
}
//int逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}
func (self *IUSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
//long左位移
type LSHL struct {
	base.NoOperandsInstruction
}
func (self *LSHL) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	s:=uint64(v2)&0x3f
	result:=int64(uint64(v1)<<s)
	stack.PushLong(result)
}
//long算数右位移
type LSHR struct {
	base.NoOperandsInstruction
}
func (self *LSHR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	s:=uint64(v2)&0x3f
	result:=v1<<s
	stack.PushLong(result)
}
//long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}
func (self *LUSHR) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	s:=uint64(v2)&0x3f
	result:=int64(uint64(v1)<<s)
	stack.PushLong(result)
}