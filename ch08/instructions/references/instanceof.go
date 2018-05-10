package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
	"jvmgo/ch07/runtime/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}
func (self *INSTANCE_OF) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	ref:=stack.PopRef()
	if ref==nil {
		stack.PushInt(0)
		return
	}
	cp:=frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class:=classRef.ResolvedClass()
	if ref.IsInstanceof(class) {
		stack.PushInt(1)
	}else {
		stack.PushInt(0)
	}

}