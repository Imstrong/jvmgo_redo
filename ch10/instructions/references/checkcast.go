package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
)

type CHECK_CAST struct {base.Index16Instruction}
func (self *CHECK_CAST) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
