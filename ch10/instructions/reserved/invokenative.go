package reserved

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/native"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}
func (self *INVOKE_NATIVE) Execute(frame *runtime.Frame) {
	method:=frame.Method()
	className:=method.Class().Name()
	methodName:=method.Name()
	methodDescriptor:=method.Descriptor()
	//根据类名、方法名、方法描述符从本地方法注册表中查找方法
	nativeMethod:=native.FindNativeMethod(className,methodName,methodDescriptor)
	if nativeMethod==nil {
		methodInfo:=className+"."+methodName+methodDescriptor
		panic("java.lang.UnsatisfiedLinkError"+methodInfo)
	}
	nativeMethod(frame)
}
