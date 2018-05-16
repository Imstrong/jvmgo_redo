package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/runtime"
	"unsafe"
)

func init() {
	native.Register("java/lang/Object","getClass","()Ljava/lang/Class;",getClass)
	native.Register("java/lang/Object","hashCode","()I",hashCode)
	native.Register("java/lang/Object","clone","()Ljava/lang/Object;",clone)
}
func getClass(frame *runtime.Frame) {
	//从局部变量表拿到this引用
	this:=frame.LocalVars().GetThis()
	//通过this引用（Object对象引用）拿到class结构体指针，通过结构体指针获取类对象
	class:=this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
//把对象引用转换成uintptr类型，在强制转换成int32推入操作数栈
func hashCode(frame *runtime.Frame) {
	this:=frame.LocalVars().GetThis()
	hash:=int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}
func clone(frame *runtime.Frame) {
	this:=frame.LocalVars().GetThis()
	cloneable:=this.Class().ClassLoader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}