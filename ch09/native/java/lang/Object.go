package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/runtime"
)

func init() {
	native.Register("java/lang/Object","getClass","()Ljava/lang/Class;",getClass)
}
func getClass(frame *runtime.Frame) {
	//从局部变量表拿到this引用
	this:=frame.LocalVars().GetThis()
	//通过this引用（Object对象引用）拿到class结构体指针，通过结构体指针获取类对象
	class:=this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
