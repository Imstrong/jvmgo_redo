package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
	"jvmgo/ch07/runtime/heap"
)
//base.Index16Instruction为带常量池索引的指令
type INVOKE_STATIC struct{base.Index16Instruction}
func (self *INVOKE_STATIC) Execute(frame *runtime.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	methodRef:=cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod:=methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeExeption")
	}
	base.InvokeMethod(frame,resolvedMethod)
}
