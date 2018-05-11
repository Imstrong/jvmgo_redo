package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
)

//计算数组长度
type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}
func (self *ARRAY_LENGTH) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	arrRef:=stack.PopRef()
	if arrRef==nil {
		panic("java.lang.NullPointerException")
	}
	arrLen:=arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
