package constants

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
	"jvmgo/ch09/runtime/heap"
)

//Load Constant
type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *runtime.Frame) {
	_ldc(frame, self.Index)
}
func (self *LDC_W) Execute(frame *runtime.Frame) {
	_ldc(frame, self.Index)
}
func _ldc(frame *runtime.Frame, index uint) {
	stack := frame.OperandStack()
	class:=frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string: //在第
		internedStr:=heap.JString(class.ClassLoader(),c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef:=c.(*heap.ClassRef)
		classObj:=classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
		panic("todo: ldc!")
	}
}
func (self *LDC2_W) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
