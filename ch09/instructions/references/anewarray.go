package references

import (
	"jvmgo_redo/ch09/instructions/base"
	"jvmgo_redo/ch09/runtime"
	"jvmgo_redo/ch09/runtime/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}
func (self *ANEW_ARRAY) Execute(frame *runtime.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	classRef:=cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass:=classRef.ResolvedClass()
	stack:=frame.OperandStack()
	count:=stack.PopInt()
	if count<0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass:=componentClass.ArrayClass()
	arr:=arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
