package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
)
//创建基本类型数组的指令
type NEW_ARRAY struct {
	//用数字表示基本类型数组的类型
	atype uint8
}
const(
	AT_BOOLEAN=4
	AT_CHAR=5
	AT_FLOAT=6
	AT_DOUBLE=7
	AT_BYTE=8
	AT_SHORT=9
	AT_INT=10
	AT_LONG=11
)
//Fetch Operands为获取操作数
func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype=reader.ReadUint8()
}
func (self *NEW_ARRAY) Execute(frame *runtime.Frame) {
	stack:=frame.OperandStack()
	count:=stack.PopInt()
	if count<0{
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader:=frame.Method().Class().ClassLoader()
	arrClass:=getPrimitiveArrayClass(classLoader,self.atype)//根据classloader，数组类型标识符加载类
	arr:=arrClass.NewArray((uint(count)))
	stack.PushRef(arr)
}
func getPrimitiveArrayClass(classLoader *heap.ClassLoader,atype uint8) *heap.Class {
	switch atype{
	case AT_BOOLEAN: return classLoader.LoadClass("[Z")
	case AT_BYTE:	return classLoader.LoadClass("[B")
	case AT_CHAR:	return classLoader.LoadClass("[C")
	case AT_SHORT:	return classLoader.LoadClass("[S")
	case AT_INT:	return classLoader.LoadClass("[I")
	case AT_LONG:	return classLoader.LoadClass("[J")
	case AT_FLOAT:	return classLoader.LoadClass("[F")
	case AT_DOUBLE:	return classLoader.LoadClass("[D")
	default:panic("Invalid array type")
	}
}

