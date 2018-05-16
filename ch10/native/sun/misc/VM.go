package misc

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
	"jvmgo/ch10/instructions/base"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}
func initialize(frame *runtime.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key:=heap.JString(vmClass.ClassLoader(),"foo")
	val:=heap.JString(vmClass.ClassLoader(),"bar")
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)
	propsClass:=vmClass.ClassLoader().LoadClass("java/util/Properties")
	setPropMethod:=propsClass.GetInstanceMethod("setProperty","(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame,setPropMethod)
}
