package lang

import (
	"jvmgo_redo/ch09/native"
	"jvmgo_redo/ch09/runtime"
	"math"
)

func init() {
	native.Register("java/lang/Float","floatToRawIntBits","(F)I",floatToRawIntBits)
}
//把float转成int
func floatToRawIntBits(frame *runtime.Frame) {
	value:=frame.LocalVars().GetFloat(0)
	bits:=math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
