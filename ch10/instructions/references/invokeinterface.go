package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
)

type INVOKE_INTERFACE struct {
	//接口方法调用指令后面是四个字节，前两字节是uint16运行时常量池索引，第三字节是给方法传递参数需要的slot数，这个数可以根据方法描述符算出来，第四字节的值是0
	index uint
}
func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index=uint(reader.ReadUint16())
	reader.ReadUint8()//slotcount
	reader.ReadUint8()//0
}
func (self *INVOKE_INTERFACE) Execute(frame *runtime.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	methodRef:=cp.GetConstant(self.index).(*heap.MethodRef)
	resolvedMethod:=methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic()||resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref:=frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
	if ref==nil {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	methodToBeInvoked:=heap.LookupMethodInClass(ref.Class(),methodRef.Name(),methodRef.Descriptor())
	if methodToBeInvoked==nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame,methodToBeInvoked)
}