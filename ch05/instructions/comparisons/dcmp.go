package comparisons

import (
	"jvmgo_redo/ch05/instructions/base"
	"jvmgo_redo/ch05/runtime"
)

//比较double变量,gflag表示操作数中是否又一个数为NaN
type DCMPG struct {
	base.NoOperandsInstruction
}
type DCMPL struct {
	base.NoOperandsInstruction
}
func _dcmp(frame *runtime.Frame,gFlag bool) {
	stack:=frame.OperandStack()
	v2:=stack.PopDouble()
	v1:=stack.PopDouble()
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
func (self *DCMPG) Execute(frame *runtime.Frame) {
	_fcmp(frame,true)
}
func (self *DCMPL) Execute(frame *runtime.Frame) {
	_fcmp(frame,false)
}
