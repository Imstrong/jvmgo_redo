package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
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
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
