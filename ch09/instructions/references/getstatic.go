package references

import (
	"jvmgo_redo/ch09/instructions/base"
	"jvmgo_redo/ch09/runtime"
	"jvmgo_redo/ch09/runtime/heap"
)

type GET_STATIC struct{base.Index16Instruction}
func (self *GET_STATIC) Execute(frame *runtime.Frame) {
	cp:=frame.Method().Class().ConstantPool()
	fieldRef:=cp.GetConstant(self.Index).(*heap.FieldRef)
	field:=fieldRef.ResolvedField()
	class:=field.Class()
	if !field.IsStatic() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(),class)
		return
	}
	descriptor:=field.Descriptor()
	slotId:=field.SlotId()
	slots:=class.StaticVars()
	stack:=frame.OperandStack()
	switch descriptor[0] {
	case 'Z','B','C','S','I':stack.PushInt(slots.GetInt(slotId))
	case 'F':stack.PushFloat(slots.GetFloat(slotId))
	case 'J':stack.PushLong(slots.GetLong(slotId))
	case 'D':stack.PushDouble(slots.GetDouble(slotId))
	case 'L','[':stack.PushRef(slots.GetRef(slotId))
	}

}
