package lang

import (
	"jvmgo_redo/ch09/native"
	"jvmgo_redo/ch09/runtime"
	"jvmgo_redo/ch09/runtime/heap"
)

func init() {
	native.Register("java/lang/String","intern","()java/lang/String;",intern)
}
func intern(frame *runtime.Frame) {
	this:=frame.LocalVars().GetThis()
	interned:=heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}