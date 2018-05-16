package lang

import (
	"jvmgo_redo/ch09/native"
	"jvmgo_redo/ch09/runtime"
	"jvmgo_redo/ch09/runtime/heap"
)

func init() {
	native.Register("java/lang/System","arrayCopy","(Ljava/lang/Object;ILjava/lang/Object;II)V",arraycopy)

}
func arraycopy(frame *runtime.Frame) {
	vars:=frame.LocalVars()
	//从局部变量表按顺序依次获得
	src:=vars.GetRef(0)
	srcPos:=vars.GetInt(1)
	dest:=vars.GetRef(2)
	destPos:=vars.GetInt(3)
	length:=vars.GetInt(4)
	if src==nil||dest==nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src,dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos<0||destPos<0||length<0 ||
		srcPos+length>src.ArrayLength()||destPos+length>dest.ArrayLength() {
			panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src,dest,srcPos,destPos,length)
}
func checkArrayCopy(src,dest *heap.Object) bool {
	srcClass:=src.Class()
	destClass:=dest.Class()
	//如果两个参数中有一个不是数组类型，
	if !srcClass.IsArray()||!destClass.IsArray(){
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive(){
			return srcClass==destClass
	}
	return true
}