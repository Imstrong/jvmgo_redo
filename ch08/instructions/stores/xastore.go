package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/runtime"
	"jvmgo/ch08/runtime/heap"
)

type AASTORE struct {base.NoOperandsInstruction}
type BASTORE struct {base.NoOperandsInstruction}
type CASTORE struct {base.NoOperandsInstruction}
type DASTORE struct {base.NoOperandsInstruction}
type IASTORE struct {base.NoOperandsInstruction}
type FASTORE struct {base.NoOperandsInstruction}
type LASTORE struct {base.NoOperandsInstruction}
type SASTORE struct {base.NoOperandsInstruction}
func checkNotNil(ref *heap.Object) {
	if ref==nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int,index int32) {
	if index<0||index>int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
func (self *IASTORE) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	val:=stack.PopInt()
	index:=stack.PopInt()
	arrRef:=stack.PopRef()
	checkNotNil(arrRef)
	ints:=arrRef.Ints()
	checkIndex(len(ints),index)
	ints[index]=int32(val)
}