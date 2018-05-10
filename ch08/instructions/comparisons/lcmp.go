package comparisons

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/runtime"
)

//比较long变量
type LCMP struct {
	base.NoOperandsInstruction
}
func (self *LCMP) Execute(frame *runtime.Frame){
	stack:=frame.OperandStack()
	v2:=stack.PopLong()
	v1:=stack.PopLong()
	if v1>v2 {
		stack.PushInt(1)
	}else if v1==v2 {
		stack.PushInt(0)
	}else {
		stack.PushInt(-1)
	}
}
