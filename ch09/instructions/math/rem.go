package math

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/runtime"
	"math"
)
//取余数指令
type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }
func (self *IREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}
//浮点数因为有Infinity（无穷大）值，因此除零也没问题
func (self *DREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
func (self *FREM) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=float64(stack.PopFloat())
	v1:=float64(stack.PopFloat())
	result:=math.Mod(v1,v2)
	stack.PushFloat(float32(result))

}
func (self *LREM) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	if v2==0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result:=v1%v2
	stack.PushLong(result)
}