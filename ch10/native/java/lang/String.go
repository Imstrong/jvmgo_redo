package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
)

func init() {
	native.Register("java/lang/String","intern","()java/lang/String;",intern)
}
func intern(frame *runtime.Frame) {
	this:=frame.LocalVars().GetThis()
	interned:=heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}