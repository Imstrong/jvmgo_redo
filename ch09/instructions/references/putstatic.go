package references

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
	"jvmgo/ch09/runtime/heap"
)

//解析静态字段
type PUT_STATIC struct{base.Index16Instruction}
func (self *PUT_STATIC) Execute(frame *runtime.Frame) {
	currentMethod:=frame.Method()
	currentClass:=currentMethod.Class()
	cp:=currentClass.ConstantPool()
	fieldRef:=cp.GetConstant(self.Index).(*heap.FieldRef)
	field:=fieldRef.ResolvedField()
	class:=field.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(),class)
		return
	}
	//如果不是静态字段，抛出异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//如果不是final字段，只能在类初始化方法中赋值
	if field.IsFinal() {
		if currentClass!=class || currentMethod.Name()!="<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	slots:=class.StaticVars()
	stack:=frame.OperandStack()
	switch descriptor[0] {
	case 'Z','B','C','S','I':slots.SetInt(slotId,stack.PopInt())
	case 'F':slots.SetFloat(slotId,stack.PopFloat())
	case 'J':slots.SetLong(slotId,stack.PopLong())
	case 'D':slots.SetDouble(slotId,stack.PopDouble())
	case 'L','[':slots.SetRef(slotId,stack.PopRef())
	}
}
