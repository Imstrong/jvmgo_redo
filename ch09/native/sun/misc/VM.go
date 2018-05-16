package misc

import (
	"jvmgo_redo/ch09/native"
	"jvmgo_redo/ch09/runtime"
	"jvmgo_redo/ch09/runtime/heap"
	"jvmgo_redo/ch09/instructions/base"
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
