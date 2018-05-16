package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/runtime"
	"jvmgo/ch10/runtime/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}
func (self *ATHROW) Execute(frame *runtime.Frame) {
	ex:=frame.OperandStack().PopRef()
	if ex==nil {
		panic("java.lang.NullPointerException")
	}
	thread:=frame.Thread()
	if !findAndGotoExcetionHandler(thread,ex) {
		handleUncaughtException(thread,ex)
	}
}
func findAndGotoExcetionHandler(thread *runtime.Thread,ref *heap.Object)bool{
	//遍历线程的栈，从当前帧开始的每一帧的异常处理表中查找该异常索引，如果找不到就弹出这一帧；找到了则把这帧的操作数栈清空，再把异常对象推入栈顶
	for{
		frame:=thread.CurrentFrame()
		pc:=frame.NextPC()-1
		handlerPC:=frame.Method().FindExceptionHandler(ref.Class(),pc)
		if handlerPC>0 {
			stack:=frame.OperandStack()
			stack.Clear()
			stack.PushRef(ref)
			frame.SetNextPC(handlerPC)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}
func handleUncaughtException(thread *runtime.Thread,ex *heap.Object) {
	thread.ClearStack()
	jMsg:=ex.GetRefVar("detailMessage","Ljava/lang/String;")
	goMsg:=heap.GoString(jMsg)
	println(ex.Class().JavaName()+":"+goMsg)
	stes:=reflect.ValueOf(ex.Extra())
	for i:=0;i<stes.len();i++ {
		ste:=stes.Index(i).Interface().(interface{String() string})
		println("\tat "+ste.String())
	}
}