package native

import "jvmgo/ch10/runtime"

//注册本地方法
type NativeMethod func(frame *runtime.Frame)
var registry=map[string]NativeMethod{}
func Register(className,methodName,methodDescriptor string,method NativeMethod) {
	key:=className+"~"+methodName+"~"+methodDescriptor
	registry[key]=method
}
func FindNativeMethod(className,methodName,methodDescriptor string) NativeMethod {
	key:=className+"~"+methodName+"~"+methodDescriptor
	if method,ok:=registry[key];ok{
		return method
	}
	if methodDescriptor=="()V"&&methodName=="registerNatives" {
		return emptyNativeMethod
	}
	return nil
}
func emptyNativeMethod(frame *runtime.Frame) {

}
