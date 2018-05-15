package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/runtime"
	"jvmgo/ch09/runtime/heap"
)

func init() {
	native.Register("java/lang/String","intern","()java/lang/String;",intern)
}
func intern(frame *runtime.Frame) {
	this:=frame.LocalVars().GetThis()
	interned:=heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}