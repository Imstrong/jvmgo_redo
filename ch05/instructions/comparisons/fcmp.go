package comparisons

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch04/runtime"
)

//比较float变量,gflag表示操作数中是否又一个数为NaN
type FCMPG struct {
	base.NoOperandsInstruction
}
type FCMPL struct {
	base.NoOperandsInstruction
}
func _fcmp(frame *runtime.Frame,gFlag bool) {
	stack:=frame.OperandStack()
	v2:=stack.PopFloat()
	v1:=stack.PopFloat()
	if v1>v2 {
		stack.PushInt(1)
	}else if v1==v2 {
		stack.PushInt(0)
	}else if v1<v2 {
		stack.PushInt(-1)
	}else if gFlag {
		stack.PushInt(1)
	}else {
		stack.PushInt(-1)
	}
}
func (self *FCMPG) Exrcute(frame *runtime.Frame) {
	_fcmp(frame,true)
}
func (self *FCMPL) Execute(frame *runtime.Frame) {
	_fcmp(frame,false)
}
